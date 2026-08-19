package api

import (
	"context"

	desc "github.com/KEKACIK/ozon-univer-golang/loms/pkg/api/loms/v1"
)

type Service interface {
	GetStocks(sku uint32) (uint64, error)
}

type Handler struct {
	desc.UnimplementedLomsServer         //
	service                      Service //
}

var _ desc.LomsServer = (*Handler)(nil)

func NewHandler(service Service) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) GetStocks(ctx context.Context, req *desc.GetStocksRequest) (*desc.GetStocksResponse, error) {
	// TODO: решить с автогенерированной валидацией
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	count, err := h.service.GetStocks(req.Sku)
	if err != nil {
		return nil, err
	}

	return &desc.GetStocksResponse{
		Count: count,
	}, nil
}
