package handlers

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/pkg"
)

type ClearHandler struct {
	name string
}

func NewClearHandler() *ClearHandler {

	return &ClearHandler{
		name: "clear handler",
	}
}

type ClearRequest struct {
	User int64 `json:"user,omitempty"`
}

func (r ClearRequest) Validate() error {

	return nil
}

func (h ClearHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !pkg.CheckMethodPost(w, r) {
		return
	}

	req := ClearRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	/* LOGIC

	cart.DeleteItemsByUserID() CartStorage

	if ok:
		return 200
	if fail:
		return error

	*/

	pkg.GetSuccessResponse(w, http.StatusOK)
}
