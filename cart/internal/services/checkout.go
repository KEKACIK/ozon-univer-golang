package service

import (
	"context"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/models"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/repository/carts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateProvider interface {
	Create(ctx context.Context, user int64, items []*models.CartModel) (int64, error)
}

type CheckoutService struct {
	dbPool         *pgxpool.Pool
	createProvider CreateProvider
}

func NewCheckoutService(
	dbPool *pgxpool.Pool,
	createProvider CreateProvider,
) *CheckoutService {

	return &CheckoutService{
		dbPool:         dbPool,
		createProvider: createProvider,
	}
}

func (s *CheckoutService) converter(items []carts.Cart) []*models.CartModel {
	result := []*models.CartModel{}
	for _, item := range items {
		result = append(result, &models.CartModel{
			UserID: int64(item.UserID),
			SKU:    uint32(item.Sku),
			Count:  uint16(item.Count),
		})
	}

	return result
}

func (s *CheckoutService) Checkout(ctx context.Context, user int64) (int64, error) {

	cartRepo := carts.New(s.dbPool)

	items, err := cartRepo.GetAllItemByUser(ctx, int32(user))
	if err != nil {
		return 0, err
	}
	cartItems := s.converter(items)

	orderId, err := s.createProvider.Create(ctx, user, cartItems)
	if err != nil {
		return 0, err
	}

	err = cartRepo.DeleteItemsByUserId(ctx, int32(user))
	if err != nil {
		return 0, err
	}

	return int64(orderId), nil
}
