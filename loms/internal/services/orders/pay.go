package orders

import (
	"errors"
	"route256/loms/internal/models"
)

type PayProvider interface {
	GetByIdOrder(orderID int64) (models.OrderModel, error)
	SetStatusOrder(orderID int64, status models.OrderStatus) error
	ReserveStock(sku uint32, count uint16) error
	UnReserveWithBuyStock(sku uint32, count uint16) error
}

type PayService struct {
	payProvider PayProvider
}

func NewPayService(payProvider PayProvider) *PayService {

	return &PayService{
		payProvider: payProvider,
	}
}

var (
	ErrOrderPay = errors.New("pay order failed")
)

func (s PayService) UnReservedAllWithFail(items []models.OrderItemModel) {
	for _, v := range items {
		_ = s.payProvider.ReserveStock(v.SKU, v.Count)
		// TODO: ignore error!!!
	}
}

func (s PayService) PayOrder(orderID int64) error {
	order, err := s.payProvider.GetByIdOrder(orderID)
	if err != nil {
		return err
	}

	reserved := []models.OrderItemModel{}
	isFail := false
	for _, v := range order.Items {
		err = s.payProvider.UnReserveWithBuyStock(v.SKU, v.Count)
		if err != nil {
			isFail = true
			break
		}
		reserved = append(reserved, v)
	}

	if isFail {
		s.UnReservedAllWithFail(reserved)
		return err
	}

	err = s.payProvider.SetStatusOrder(orderID, models.OrderStatusPayed)
	if err != nil {
		s.UnReservedAllWithFail(reserved)
		return err
	}

	return nil
}
