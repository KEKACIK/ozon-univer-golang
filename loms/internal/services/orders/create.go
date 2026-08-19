package orders

import (
	"errors"

	"github.com/KEKACIK/ozon-univer-golang/loms/internal/models"
)

type CreateProvider interface {
	CreateOrder(userID int64, items []models.OrderItemModel, status models.OrderStatus) (int64, error)
	SetStatusOrder(orderID int64, status models.OrderStatus) error
	ReserveStock(sku uint32, count uint16) error
	UnReserveStock(sku uint32, count uint16) error
}

type CreateService struct {
	createProvider CreateProvider
}

func NewCreateService(createProvider CreateProvider) *CreateService {

	return &CreateService{
		createProvider: createProvider,
	}
}

var (
	ErrOrderCreate = errors.New("create order failed")
)

func (s CreateService) Create(userID int64, items []models.OrderItemModel) (int64, error) {
	orderID, err := s.createProvider.CreateOrder(userID, items, models.OrderStatusNew)
	if err != nil {
		return 0, err
	}

	reservedItems := []models.OrderItemModel{}
	isFail := false
	for _, oItem := range items {
		err = s.createProvider.ReserveStock(oItem.SKU, oItem.Count)
		if err != nil {
			isFail = true
			break
		}
		reservedItems = append(reservedItems, oItem)
	}

	if isFail {
		for _, rItem := range reservedItems {
			// TODO: ignore error!!!
			_ = s.createProvider.UnReserveStock(rItem.SKU, rItem.Count)
		}
		s.createProvider.SetStatusOrder(orderID, models.OrderStatusFailed)
		return 0, ErrOrderCreate
	}
	s.createProvider.SetStatusOrder(orderID, models.OrderStatusWaitPayment)

	return orderID, nil
}
