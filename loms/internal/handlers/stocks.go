package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"route256/loms/internal/pkg"
)

type StocksService interface {
	GetStocks(sku uint32) uint64
}

type StocksHandler struct {
	handlerName string

	stockService StocksService
}

func NewStocksHandler(stockService StocksService) *StocksHandler {

	return &StocksHandler{
		handlerName:  "stocks handler",
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
	req := &StocksRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		pkg.GetErrorResponse(w, h.handlerName, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		pkg.GetErrorResponse(w, h.handlerName, err, http.StatusBadRequest)
		return
	}

	count := h.stockService.GetStocks(req.SKU)

	stocksResponse := StocksResponse{
		Count: count,
	}

	raw, err := json.Marshal(stocksResponse)
	if err != nil {
		pkg.GetErrorResponse(w, h.handlerName, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	pkg.GetSuccessResponseWithBody(w, raw, http.StatusOK)
}
