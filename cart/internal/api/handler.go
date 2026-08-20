package api

import (
	desc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/cart/v1"
)

type Handler struct {
	desc.UnimplementedCartServer

	itemADder ItemAdder
}

var _ desc.CartServer = (*Handler)(nil)

func NewHandler(
	itemADder ItemAdder,
) *Handler {

	return &Handler{
		itemADder: itemADder,
	}
}
