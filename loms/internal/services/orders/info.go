package orders

import (
	"route256/loms/internal/models"
)

type InfoProvider interface {
	GetByIdOrder(orderID int64) (models.OrderModel, error)
}

type InfoService struct {
	infoProvider InfoProvider
}

func NewInfoService(infoProvider InfoProvider) *InfoService {

	return &InfoService{
		infoProvider: infoProvider,
	}
}

func (s InfoService) InfoOrder(orderID int64) (models.OrderModel, error) {
	order, err := s.infoProvider.GetByIdOrder(orderID)
	if err != nil {
		return models.OrderModel{}, err
	}

	return order, nil
}
