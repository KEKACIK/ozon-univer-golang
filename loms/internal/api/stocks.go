package api

import (
	"context"

	desc "github.com/KEKACIK/ozon-univer-golang/loms/pkg/api/loms/v1"
)

type StockReader interface {
	GetStocks(sku uint32) (uint64, error)
}

func (h *Handler) GetStocks(ctx context.Context, req *desc.GetStocksRequest) (*desc.GetStocksResponse, error) {
	// TODO: решить с автогенерированной валидацией
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	count, err := h.stockReader.GetStocks(req.Sku)
	if err != nil {
		return nil, err
	}

	return &desc.GetStocksResponse{
		Count: count,
	}, nil
}
