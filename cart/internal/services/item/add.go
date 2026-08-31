package item

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/repository/carts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StocksProvider interface {
	GetStocks(ctx context.Context, sku uint32) (uint64, error)
}

type ProductProvider interface {
	GetProductInfo(ctx context.Context, sku uint32) (string, uint32, error)
}

type AddService struct {
	name   string
	dbPool *pgxpool.Pool

	stocksProvider  StocksProvider
	productProvider ProductProvider
}

func NewAddService(
	stocksProvider StocksProvider,
	productProvider ProductProvider,
	dbPool *pgxpool.Pool,
) *AddService {

	return &AddService{
		name:   "item add service",
		dbPool: dbPool,

		stocksProvider:  stocksProvider,
		productProvider: productProvider,
	}
}

var (
	ErrInsufficientStocks = errors.New("insufficient stocks")
)

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
		return fmt.Errorf("%s: %w", s.name, ErrInsufficientStocks)
	}

	cartRepo := carts.New(s.dbPool)
	_, err = cartRepo.CreateItem(
		ctx,
		carts.CreateItemParams{},
	)
	if err != nil {
		return err
	}

	return nil
}
