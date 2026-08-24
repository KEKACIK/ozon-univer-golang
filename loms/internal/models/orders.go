package models

const (
	OrderStatusNew         string = "new"
	OrderStatusWaitPayment string = "awaiting payment"
	OrderStatusFailed      string = "failed"
	OrderStatusPayed       string = "payed"
	OrderStatusCanceled    string = "canceled"
)

type OrderItemModel struct {
	SKU   uint32
	Count uint16
}

type OrderModel struct {
	ID     int64
	UserID int64
	Status string
	Items  []OrderItemModel
}
