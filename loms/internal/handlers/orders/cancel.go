package orders

import (
	"encoding/json"
	"net/http"
	"route256/loms/internal/pkg"
)

type CancelHandler struct {
	name string
}

func NewCancelHandler() *CancelHandler {

	return &CancelHandler{
		name: "orders cancel handler",
	}
}

type CancelRequest struct {
	OrderId int64
}

func (r CancelRequest) Validate() error {

	return nil
}

func (h CancelHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &CancelRequest{}
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
	stocks.ReserveCancel() StocksStorage (отменены)
	order.SetStatus(cancelled) OrderStorage

	if ok:
		return 200
	if failt:
		return error (not found)

	*/

	pkg.GetSuccessResponse(w, http.StatusOK)
}
