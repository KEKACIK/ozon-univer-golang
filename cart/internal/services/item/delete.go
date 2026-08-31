package item

import (
	"context"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/repository/carts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DeleteService struct {
	dbPool          *pgxpool.Pool
	productProvider ProductProvider
}

func NewDeleteService(
	dbPool *pgxpool.Pool,
	productProvider ProductProvider,
) *DeleteService {

	return &DeleteService{
		dbPool:          dbPool,
		productProvider: productProvider,
	}
}

func (s *DeleteService) Delete(
	ctx context.Context,
	user int64,
	sku uint32,
) error {
	if _, _, err := s.productProvider.GetProductInfo(ctx, sku); err != nil {
		return err
	}

	cartRepo := carts.New(s.dbPool)
	err := cartRepo.DeleteItemBySku(
		ctx,
		carts.DeleteItemBySkuParams{
			UserID: int32(user),
			Sku:    int32(sku),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
