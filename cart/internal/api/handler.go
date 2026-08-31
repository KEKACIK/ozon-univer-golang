package api

import (
	desc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/cart/v1"
)

type Handler struct {
	desc.UnimplementedCartServiceServer

	cartLister     CartLister
	cartClearer    CartClearer
	cartCheckouter CartCheckouter
	itemAdder      ItemAdder
	itemDeleter    ItemDeleter
}

var _ desc.CartServiceServer = (*Handler)(nil)

func NewHandler(
	cartLister CartLister,
	cartClearer CartClearer,
	cartCheckouter CartCheckouter,
	itemAdder ItemAdder,
	itemDeleter ItemDeleter,
) *Handler {

	return &Handler{
		cartLister:     cartLister,
		cartClearer:    cartClearer,
		cartCheckouter: cartCheckouter,

		itemAdder:   itemAdder,
		itemDeleter: itemDeleter,
	}
}
