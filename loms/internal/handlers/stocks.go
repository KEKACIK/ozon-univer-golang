package handlers

import (
	"errors"
	"net/http"

	desc "github.com/KEKACIK/ozon-univer-golang/loms/pkg/api/loms/v1"
)

type LomsServer struct {
	desc.UnimplementedLomsServer
}

type StocksService interface {
	GetStocks(sku uint32) (uint64, error)
}

type StocksHandler struct {
	name string

	stockService StocksService
}

func NewStocksHandler(stockService StocksService) *StocksHandler {

	return &StocksHandler{
		name:         "stocks handler",
		stockService: stockService,
	}
}

type StocksRequest struct {
	SKU uint32 `json:"sku,omitempty"`
}

var (
	ErrIncorrectSKU = errors.New("incorrect SKU")
)

func (r StocksRequest) Validate() error {
	if r.SKU == 0 {
		return ErrIncorrectSKU
	}

	return nil
}

type StocksResponse struct {
	Count uint64 `json:"count,omitempty"`
}

func (h StocksHandler) Handle(w http.ResponseWriter, r *http.Request) {

}
