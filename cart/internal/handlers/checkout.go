package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/pkg"
)

type CheckoutHandler struct {
	name string
}

func NewCheckoutHandler() *CheckoutHandler {

	return &CheckoutHandler{
		name: "checkout handler",
	}
}

type CheckoutRequest struct {
	User int64 `json:"user,omitempty"`
}

func (r CheckoutRequest) Validate() error {

	return nil
}

func (h CheckoutHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !pkg.CheckMethodPost(w, r) {
		return
	}

	req := CheckoutRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	/* LOGIC

	cart.GetItemsByUserID() CartStorage
	orders/create(user, []items) Loms/order
	cart.DeleteItemsByUserId CartStorage

	if ok:
		return 200
	if fail:
		return error

	*/

	pkg.GetSuccessResponse(w, http.StatusOK)
}
