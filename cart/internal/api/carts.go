package api

import (
	"context"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/api/converter"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/models"
	desc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/cart/v1"
)

type CartLister interface {
	List(context.Context, int64) ([]*models.CartModel, uint32, error)
}

func (h *Handler) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {

	items, sum, err := h.cartLister.List(ctx, req.User)
	if err != nil {
		return nil, err
	}

	return converter.CartListConvertModel2Response(items, sum), nil
}
