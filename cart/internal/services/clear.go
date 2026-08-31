package service

import (
	"context"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/repository/carts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClearService struct {
	dbPool          *pgxpool.Pool
	productProvider ProductProvider
}

func NewClearService(
	dbPool *pgxpool.Pool,
) *ClearService {

	return &ClearService{
		dbPool: dbPool,
	}
}
func (s *ClearService) Clear(ctx context.Context, user int64) error {
	cartRepo := carts.New(s.dbPool)

	err := cartRepo.DeleteItemsByUserId(ctx, int32(user))
	if err != nil {
		return err
	}

	return nil
}
