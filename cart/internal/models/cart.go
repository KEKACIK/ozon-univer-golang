package models

type CartModel struct {
	UserID int64
	SKU    uint32
	Count  uint16
	Name   string
	Price  uint32
}
