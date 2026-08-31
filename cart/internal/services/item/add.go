package item

import (
	"context"
	"errors"
	"time"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/repository/carts"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StocksProvider interface {
	GetStocks(ctx context.Context, sku uint32) (uint64, error)
}

type ProductProvider interface {
	GetProductInfo(ctx context.Context, sku uint32) (string, uint32, error)
}

type AddService struct {
	dbPool *pgxpool.Pool

	stocksProvider  StocksProvider
	productProvider ProductProvider
}

func NewAddService(
	dbPool *pgxpool.Pool,
	stocksProvider StocksProvider,
	productProvider ProductProvider,
) *AddService {

	return &AddService{
		dbPool:          dbPool,
		stocksProvider:  stocksProvider,
		productProvider: productProvider,
	}
}

var (
	ErrInsufficientStocks = errors.New("insufficient stocks")
)

func (s *AddService) CreateOrUpdateItem(
	ctx context.Context,
	cartRepo *carts.Queries,
	user int64,
	sku uint32,
	count uint16,
) error {
	cart, err := cartRepo.GetItemBySku(
		ctx,
		carts.GetItemBySkuParams{
			UserID: int32(user),
			Sku:    int32(sku),
		},
	)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}

		_, err = cartRepo.CreateItem(
			ctx,
			carts.CreateItemParams{
				UserID: int32(user),
				Sku:    int32(sku),
				Count:  int32(count),
			},
		)

		return err
	}

	err = cartRepo.UpdateItemCount(
		ctx,
		carts.UpdateItemCountParams{
			ID:    cart.ID,
			Count: int32(count),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *AddService) Add(
	ctx context.Context,
	user int64,
	sku uint32,
	count uint16,
) error {
	if _, _, err := s.productProvider.GetProductInfo(ctx, sku); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	stockCount, err := s.stocksProvider.GetStocks(ctx, sku)
	if err != nil {
		return err
	}

	if uint64(count) > stockCount {
		return ErrInsufficientStocks
	}

	cartRepo := carts.New(s.dbPool)

	err = s.CreateOrUpdateItem(ctx, cartRepo, user, sku, count)

	if err != nil {
		return err
	}

	return nil
}
