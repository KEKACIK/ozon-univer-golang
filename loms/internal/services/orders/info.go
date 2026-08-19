package orders

import (
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/models"
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

func (s InfoService) GetInfo(orderID int64) (models.OrderModel, error) {
	order, err := s.infoProvider.GetByIdOrder(orderID)
	if err != nil {
		return models.OrderModel{}, err
	}

	return order, nil
}
