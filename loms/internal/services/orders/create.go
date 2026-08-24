package orders

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/KEKACIK/ozon-univer-golang/loms/internal/models"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/repository/orders"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/repository/orders_items"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/repository/stocks"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateService struct {
	dbPool *pgxpool.Pool
}

func NewCreateService(dbPool *pgxpool.Pool) *CreateService {

	return &CreateService{dbPool: dbPool}
}

var (
	ErrOrderCreate         = errors.New("create order failed")
	ErrOrderStockNotEnough = errors.New("Stock.%d not enough")
)

func (s *CreateService) Create(ctx context.Context, userID int64, items []*models.OrderItemModel) (int64, error) {
	orderRepo := orders.New(s.dbPool)

	order, err := orderRepo.CreateOrder(
		ctx,
		orders.CreateOrderParams{
			UserID: int32(userID),
			Status: models.OrderStatusNew,
		},
	)
	if err != nil {
		return 0, err
	}

	failFunc := func() {
		orderRepo.SetStatusOrder(
			ctx,
			orders.SetStatusOrderParams{
				ID:     order.ID,
				Status: models.OrderStatusFailed,
			},
		)
	}

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		failFunc()
		return 0, err
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			if errors.Is(err, pgx.ErrTxClosed) {
				return
			}
			log.Panicln(fmt.Errorf("rollback transaction: %w", err))
		}
	}()

	orderTxRepo := orderRepo.WithTx(tx)

	orderItemRepo := orders_items.New(s.dbPool)
	orderItemRepo = orderItemRepo.WithTx(tx)

	stockRepo := stocks.New(s.dbPool)
	stockRepo = stockRepo.WithTx(tx)

	for _, item := range items {
		stock, err := stockRepo.GetStocks(ctx, int32(item.SKU))
		if err != nil {
			failFunc()
			return 0, err
		}
		stockFree := stock.Count - stock.Reserved
		if stockFree < int32(item.Count) {
			failFunc()
			return 0, fmt.Errorf(ErrOrderStockNotEnough.Error(), item.SKU)
		}

		err = stockRepo.UpdateStock(
			ctx,
			stocks.UpdateStockParams{
				Sku:      int32(item.SKU),
				Count:    stock.Count,
				Reserved: stock.Reserved + int32(item.Count),
			},
		)
		if err != nil {
			failFunc()
			return 0, err
		}

		_, err = orderItemRepo.CreateOrderItem(
			ctx,
			orders_items.CreateOrderItemParams{
				OrderID: int64(order.ID),
				Sku:     int32(item.SKU),
				Count:   int32(item.Count),
			},
		)
		if err != nil {
			failFunc()
			return 0, err
		}
	}

	orderTxRepo.SetStatusOrder(
		ctx,
		orders.SetStatusOrderParams{
			ID:     order.ID,
			Status: models.OrderStatusWaitPayment,
		},
	)

	if err := tx.Commit(ctx); err != nil {
		failFunc()
		return 0, err
	}

	// orderID, err := s.createProvider.CreateOrder(userID, items, models.OrderStatusNew)
	// if err != nil {
	// 	return 0, err
	// }

	// reservedItems := []models.OrderItemModel{}
	// isFail := false
	// for _, oItem := range items {
	// 	err = s.createProvider.ReserveStock(oItem.SKU, oItem.Count)
	// 	if err != nil {
	// 		isFail = true
	// 		break
	// 	}
	// 	reservedItems = append(reservedItems, oItem)
	// }

	// if isFail {
	// 	for _, rItem := range reservedItems {
	// 		// TODO: ignore error!!!
	// 		_ = s.createProvider.UnReserveStock(rItem.SKU, rItem.Count)
	// 	}
	// 	s.createProvider.SetStatusOrder(orderID, models.OrderStatusFailed)
	// 	return 0, ErrOrderCreate
	// }
	// s.createProvider.SetStatusOrder(orderID, models.OrderStatusWaitPayment)

	return int64(order.ID), nil
}
