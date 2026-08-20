package api

import (
	"context"

	desc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/cart/v1"
)

type ItemAdder interface {
	Add(ctx context.Context, User int64, SKU uint32, Count uint16) error
}

func (h *Handler) ItemAdd(ctx context.Context, req *desc.ItemAddRequest) (*desc.ItemAddResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	err := h.itemADder.Add(ctx, req.User, req.Sku, uint16(req.Count))
	if err != nil {
		return nil, err
	}

	return &desc.ItemAddResponse{}, nil
}
