package api

import (
	"context"

	"github.com/KEKACIK/ozon-univer-golang/loms/internal/api/converter"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/models"
	desc "github.com/KEKACIK/ozon-univer-golang/loms/pkg/api/loms/v1"
)

type OrderCreator interface {
	Create(ctx context.Context, userID int64, items []*models.OrderItemModel) (int64, error)
}

type OrderInfoReader interface {
	GetInfo(ctx context.Context, orderID int64) (*models.OrderModel, error)
}

type OrderPayer interface {
	Pay(ctx context.Context, orderID int64) error
}

type OrderCanceller interface {
	Cancel(ctx context.Context, orderID int64) error
}

func (h *Handler) OrderCreate(ctx context.Context, req *desc.OrderCreateRequest) (*desc.OrderCreateResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	orderID, err := h.orderCreator.Create(ctx, req.User, converter.OrderConvertCreateRequest2Model(req))
	if err != nil {
		return nil, err
	}

	return &desc.OrderCreateResponse{
		OrderId: orderID,
	}, nil
}

func (h *Handler) OrderInfo(ctx context.Context, req *desc.OrderInfoRequest) (*desc.OrderInfoResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	order, err := h.orderInfoReader.GetInfo(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	return converter.OrderConvertModel2InfoResponse(order), nil
}

func (h *Handler) OrderPay(ctx context.Context, req *desc.OrderPayRequest) (*desc.OrderPayResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	err := h.orderPayer.Pay(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &desc.OrderPayResponse{}, nil
}

func (h *Handler) OrderCancel(ctx context.Context, req *desc.OrderCancelRequest) (*desc.OrderCancelResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	err := h.orderCanceller.Cancel(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &desc.OrderCancelResponse{}, nil
}
