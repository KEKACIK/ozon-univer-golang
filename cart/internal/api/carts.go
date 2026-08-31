package api

import (
	"context"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/api/converter"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/models"
	desc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/cart/v1"
)

type CartLister interface {
	List(ctx context.Context, user int64) ([]*models.CartModel, uint32, error)
}

type CartClearer interface {
	Clear(ctx context.Context, user int64) error
}

type CartCheckouter interface {
	Checkout(ctx context.Context, user int64) (int64, error)
}

func (h *Handler) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {

	items, sum, err := h.cartLister.List(ctx, req.User)
	if err != nil {
		return nil, err
	}

	return converter.CartListConvertModel2Response(items, sum), nil
}

func (h *Handler) Clear(ctx context.Context, req *desc.ClearRequest) (*desc.ClearResponse, error) {
	err := h.cartClearer.Clear(ctx, req.User)
	if err != nil {
		return nil, err
	}

	return &desc.ClearResponse{}, nil
}

func (h *Handler) Checkout(ctx context.Context, req *desc.CheckoutRequest) (*desc.CheckoutResponse, error) {

	orderId, err := h.cartCheckouter.Checkout(ctx, req.User)
	if err != nil {
		return nil, err
	}

	return &desc.CheckoutResponse{OrderId: orderId}, nil
}
