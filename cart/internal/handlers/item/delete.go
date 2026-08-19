package item

import (
	"encoding/json"
	"net/http"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/pkg"
)

type DeleteHandler struct {
	name string
}

func NewDeleteHandler() *DeleteHandler {

	return &DeleteHandler{
		name: "item delete handler",
	}
}

type DeleteRequest struct {
	User int64
	SKU  uint32
}

func (r DeleteRequest) Validate() error {

	return nil
}

func (h DeleteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !pkg.CheckMethodPost(w, r) {
		return
	}

	req := DeleteRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	/* LOGIC

	cart.DeleteItem() CartStorage

	if ok:
		return 200
	if fail:
		return error

	*/

	pkg.GetSuccessResponse(w, http.StatusOK)
}
