package orders

import (
	"encoding/json"
	"net/http"
	"route256/loms/internal/pkg"
)

type CreateHandler struct {
	name string
}

func NewCreateHandler() *CreateHandler {

	return &CreateHandler{
		name: "orders create handler",
	}
}

type CreateItemRequest struct {
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

type CreateRequest struct {
	User  int64
	items []CreateItemRequest
}

func (r CreateRequest) Validate() error {

	return nil
}

type CreateResponse struct {
	OrderId uint64 `json:"order_id,omitempty"`
}

func (h CreateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	/* LOGIC

	order.Create() OrderStorage
	order.SetStatus(new) OrderStorage
	stocks.Reserve() StocksStorage

	if ok:
		order.SetStatus(awaiting_payment) OrderStorage
		return response
	if failt:
		order.SetStatus(failed) OrderStorage
		return error

	*/

	createResponse := CreateResponse{
		OrderId: 1,
	}

	raw, err := json.Marshal(createResponse)
	if err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	pkg.GetSuccessResponseWithBody(w, raw, http.StatusOK)
}
