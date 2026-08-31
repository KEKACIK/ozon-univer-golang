package api

import (
	desc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/cart/v1"
)

type Handler struct {
	desc.UnimplementedCartServiceServer

	itemAdder ItemAdder
}

var _ desc.CartServiceServer = (*Handler)(nil)

func NewHandler(
	itemAdder ItemAdder,
) *Handler {

	return &Handler{
		itemAdder: itemAdder,
	}
}
