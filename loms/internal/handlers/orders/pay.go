package orders

import (
	"encoding/json"
	"net/http"
	"route256/loms/internal/pkg"
)

type PayService interface {
	PayOrder(orderID int64) error
}

type PayHandler struct {
	name       string
	payService PayService
}

func NewPayHandler(payService PayService) *PayHandler {

	return &PayHandler{
		name:       "orders pay handler",
		payService: payService,
	}
}

type PayRequest struct {
	OrderId int64 `json:"order_id,omitempty"`
}

func (r PayRequest) Validate() error {

	return nil
}

func (h PayHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !pkg.CheckMethodPost(w, r) {
		return
	}

	req := &PayRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	err := h.payService.PayOrder(req.OrderId)

	if err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	pkg.GetSuccessResponse(w, http.StatusOK)
}
