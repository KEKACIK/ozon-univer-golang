package orders

import (
	"route256/loms/internal/models"
)

type CancelProvider interface {
	GetByIdOrder(orderID int64) (models.OrderModel, error)
	SetStatusOrder(orderID int64, status models.OrderStatus) error
	UnReserveStock(sku uint32, count uint16) error
}

type CancelService struct {
	cancelProvider CancelProvider
}

func NewCancelService(cancelProvider CancelProvider) *CancelService {

	return &CancelService{
		cancelProvider: cancelProvider,
	}
}

func (s CancelService) CancelOrder(orderID int64) error {
	order, err := s.cancelProvider.GetByIdOrder(orderID)
	if err != nil {
		return err
	}

	for _, v := range order.Items {
		_ = s.cancelProvider.UnReserveStock(v.SKU, v.Count)
		// TODO: ignored error!!!
	}
	err = s.cancelProvider.SetStatusOrder(orderID, models.OrderStatusCanceled)

	return nil
}
