package orders

import (
	"context"

	"github.com/KEKACIK/ozon-univer-golang/loms/internal/models"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/repository/orders"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/repository/orders_items"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InfoService struct {
	dbPool *pgxpool.Pool
}

func NewInfoService(dbPool *pgxpool.Pool) *InfoService {

	return &InfoService{dbPool: dbPool}
}

func ConvertItems(items []orders_items.OrdersItem) ([]models.OrderItemModel, error) {
	result := make([]models.OrderItemModel, len(items))
	for i, v := range items {
		result[i] = models.OrderItemModel{
			SKU:   uint32(v.Sku),
			Count: uint16(v.Count),
		}
	}

	return result, nil
}

func (s *InfoService) GetInfo(ctx context.Context, orderID int64) (*models.OrderModel, error) {
	orderRepo := orders.New(s.dbPool)
	orderItemsRepo := orders_items.New(s.dbPool)

	order, err := orderRepo.GetByIdOrder(ctx, int32(orderID))
	if err != nil {
		return nil, err
	}

	orderItems, err := orderItemsRepo.GetByOrderId(ctx, orderID)
	if err != nil {
		return nil, err
	}

	items, err := ConvertItems(orderItems)
	if err != nil {
		return nil, err
	}

	return &models.OrderModel{
		ID:     int64(order.ID),
		UserID: int64(order.UserID),
		Status: order.Status,
		Items:  items,
	}, nil
}
