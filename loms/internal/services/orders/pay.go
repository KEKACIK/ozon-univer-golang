package orders

import (
	"context"
	"errors"

	"github.com/KEKACIK/ozon-univer-golang/loms/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PayService struct {
	dbPool *pgxpool.Pool
}

func NewPayService(dbPool *pgxpool.Pool) *PayService {

	return &PayService{dbPool: dbPool}
}

var (
	ErrOrderPay = errors.New("pay order failed")
)

func (s PayService) UnReservedAllWithFail(items []models.OrderItemModel) {
	// for _, v := range items {
	// 	_ = s.payProvider.ReserveStock(v.SKU, v.Count)
	// 	// TODO: ignore error!!!
	// }
}

func (s PayService) Pay(ctx context.Context, orderID int64) error {
	// order, err := s.payProvider.GetByIdOrder(orderID)
	// if err != nil {
	// 	return err
	// }

	// reserved := []models.OrderItemModel{}
	// isFail := false
	// for _, v := range order.Items {
	// 	err = s.payProvider.UnReserveWithBuyStock(v.SKU, v.Count)
	// 	if err != nil {
	// 		isFail = true
	// 		break
	// 	}
	// 	reserved = append(reserved, v)
	// }

	// if isFail {
	// 	s.UnReservedAllWithFail(reserved)
	// 	return err
	// }

	// err = s.payProvider.SetStatusOrder(orderID, models.OrderStatusPayed)
	// if err != nil {
	// 	s.UnReservedAllWithFail(reserved)
	// 	return err
	// }

	return nil
}
