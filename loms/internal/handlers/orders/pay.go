package orders

import (
	"encoding/json"
	"net/http"
	"route256/loms/internal/pkg"
)

type PayHandler struct {
	name string
}

func NewPayHandler() *PayHandler {

	return &PayHandler{
		name: "orders pay handler",
	}
}

type PayRequest struct {
	OrderId int64
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

	/* LOGIC

	order.GetByID() OrderStorage
	stocks.ReserveBuy() StocksStorage (куплены)
	order.SetStatus(payed) OrderStorage

	if ok:
		return 200
	if failt:
		return error (not found)

	*/

	pkg.GetSuccessResponse(w, http.StatusOK)
}
