package models

type OrderStatus string

const (
	OrderStatusNew         OrderStatus = "new"
	OrderStatusWaitPayment OrderStatus = "awaiting payment"
	OrderStatusFailed      OrderStatus = "failed"
	OrderStatusPayed       OrderStatus = "payed"
	OrderStatusCanceled    OrderStatus = "canceled"
)

type OrderItemModel struct {
	SKU   uint32
	Count uint16
}

type OrderModel struct {
	ID     int64
	UserID int64
	Status OrderStatus
	Items  []OrderItemModel
}
