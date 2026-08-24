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

type PayService struct {
	dbPool *pgxpool.Pool
}

func NewPayService(dbPool *pgxpool.Pool) *PayService {

	return &PayService{dbPool: dbPool}
}

var (
	ErrOrderPayStatusInvalid = errors.New("pay order failed: status must be '%s'")
	ErrOrderPayStockInvalid  = errors.New("pay order failed: stock.%d invalid")
)

func (s *PayService) Pay(ctx context.Context, orderID int64) error {
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
	if order.Status != models.OrderStatusWaitPayment {
		return fmt.Errorf(ErrOrderPayStatusInvalid.Error(), models.OrderStatusWaitPayment)
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

		newCount := stock.Count - item.Count
		newReserved := stock.Reserved - item.Count
		if newCount < 0 || newReserved < 0 {
			return fmt.Errorf(ErrOrderPayStockInvalid.Error(), stock.Sku)
		}

		err = stockRepo.UpdateStock(
			ctx,
			stocks.UpdateStockParams{
				Sku:      stock.Sku,
				Count:    newCount,
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
			Status: models.OrderStatusPayed,
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
