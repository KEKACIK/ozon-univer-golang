package api

import (
	desc "github.com/KEKACIK/ozon-univer-golang/loms/pkg/api/loms/v1"
)

type Handler struct {
	desc.UnimplementedLomsServer

	orderCreator    OrderCreator
	orderInfoReader OrderInfoReader
	orderPayer      OrderPayer
	orderCanceller  OrderCanceller

	stockReader StockReader
}

var _ desc.LomsServer = (*Handler)(nil)

func NewHandler(
	orderCreator OrderCreator,
	orderInfoReader OrderInfoReader,
	orderPayer OrderPayer,
	orderCanceller OrderCanceller,
	stockReader StockReader,
) *Handler {

	return &Handler{
		orderCreator:    orderCreator,
		orderInfoReader: orderInfoReader,
		orderPayer:      orderPayer,
		orderCanceller:  orderCanceller,
		stockReader:     stockReader,
	}
}
