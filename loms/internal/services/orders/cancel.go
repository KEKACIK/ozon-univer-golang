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

type CancelService struct {
	dbPool *pgxpool.Pool
}

func NewCancelService(dbPool *pgxpool.Pool) *CancelService {

	return &CancelService{dbPool: dbPool}
}

var (
	ErrInvalidStockReserved = errors.New("invalid stock.%d reserve value")
)

func (s *CancelService) Cancel(ctx context.Context, orderID int64) error {
	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			if errors.Is(err, pgx.ErrTxClosed) {
				return
			}
			log.Panicln(fmt.Errorf("rollback transaction: %w", err))
		}
	}()

	orderRepo := orders.New(s.dbPool)
	orderRepo = orderRepo.WithTx(tx)

	order, err := orderRepo.GetByIdOrder(ctx, int32(orderID))
	if err != nil {
		return err
	}

	orderItemRepo := orders_items.New(s.dbPool)
	orderItemRepo = orderItemRepo.WithTx(tx)

	orderItems, err := orderItemRepo.GetByOrderId(ctx, int64(order.ID))
	if err != nil {
		return err
	}

	stockRepo := stocks.New(s.dbPool)
	stockRepo = stockRepo.WithTx(tx)

	for _, item := range orderItems {
		stock, err := stockRepo.GetStocks(ctx, item.Sku)
		if err != nil {
			return err
		}
		newReserved := stock.Reserved - item.Count
		if newReserved < 0 {
			return fmt.Errorf(ErrInvalidStockReserved.Error(), stock.Sku)
		}

		err = stockRepo.UpdateStock(
			ctx,
			stocks.UpdateStockParams{
				Sku:      item.Sku,
				Count:    stock.Count,
				Reserved: newReserved,
			},
		)
		if err != nil {
			return err
		}
	}

	err = orderRepo.SetStatusOrder(
		ctx,
		orders.SetStatusOrderParams{
			ID:     order.ID,
			Status: models.OrderStatusCanceled,
		},
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
