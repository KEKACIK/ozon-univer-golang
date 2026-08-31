package service

import (
	"context"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/models"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/repository/carts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductProvider interface {
	GetProductInfo(ctx context.Context, sku uint32) (string, uint32, error)
}

type ListService struct {
	dbPool          *pgxpool.Pool
	productProvider ProductProvider
}

func NewListService(
	dbPool *pgxpool.Pool,
	productProvider ProductProvider,
) *ListService {

	return &ListService{
		dbPool:          dbPool,
		productProvider: productProvider,
	}
}

func (s *ListService) converter(
	ctx context.Context,
	items []carts.Cart,
) ([]*models.CartModel, uint32, error) {
	result := []*models.CartModel{}
	sumPrice := uint32(0)

	for _, item := range items {
		name, price, err := s.productProvider.GetProductInfo(ctx, uint32(item.Sku))
		if err != nil {
			return nil, 0, err
		}

		result = append(result, &models.CartModel{
			UserID: int64(item.UserID),
			SKU:    uint32(item.Sku),
			Count:  uint16(item.Count),
			Name:   name,
			Price:  price,
		})
		sumPrice += price * uint32(item.Count)
	}

	return result, sumPrice, nil
}

func (s *ListService) List(
	ctx context.Context,
	user int64,
) ([]*models.CartModel, uint32, error) {
	cartRepo := carts.New(s.dbPool)
	items, err := cartRepo.GetAllItemByUser(ctx, int32(user))
	if err != nil {
		return nil, 0, err
	}

	cartItems, cartSum, err := s.converter(ctx, items)

	return cartItems, cartSum, nil
}
