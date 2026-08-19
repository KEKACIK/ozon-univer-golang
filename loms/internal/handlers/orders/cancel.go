package orders

import (
	"encoding/json"
	"net/http"

	"github.com/KEKACIK/ozon-univer-golang/loms/internal/pkg"
)

type CancelService interface {
	CancelOrder(orderID int64) error
}

type CancelHandler struct {
	name          string
	cancelService CancelService
}

func NewCancelHandler(cancelService CancelService) *CancelHandler {

	return &CancelHandler{
		name:          "orders cancel handler",
		cancelService: cancelService,
	}
}

type CancelRequest struct {
	OrderId int64 `json:"order_id,omitempty"`
}

func (r CancelRequest) Validate() error {
	if r.OrderId <= 0 {
		return ErrOrderIDIncorrect
	}

	return nil
}

func (h CancelHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !pkg.CheckMethodPost(w, r) {
		return
	}

	req := &CancelRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	err := h.cancelService.CancelOrder(req.OrderId)
	if err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	pkg.GetSuccessResponse(w, http.StatusOK)
}
