package repository

import (
	"errors"
	"route256/loms/internal/models"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

func (r *DumpRepo) CreateOrder(userID int64, items []models.OrderItemModel, status models.OrderStatus) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	orderID := int64(len(r.orders) + 1)
	r.orders[orderID] = models.OrderModel{
		ID:     orderID,
		UserID: userID,
		Items:  items,
		Status: models.OrderStatusNew,
	}

	return orderID, nil
}

func (r *DumpRepo) SetStatusOrder(orderID int64, status models.OrderStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, ok := r.orders[orderID]
	if !ok {
		return ErrOrderNotFound
	}
	order.Status = status
	r.orders[orderID] = order

	return nil
}

func (r *DumpRepo) GetByIdOrder(orderID int64) (models.OrderModel, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, ok := r.orders[orderID]
	if !ok {
		return models.OrderModel{}, ErrOrderNotFound
	}

	return order, nil
}
