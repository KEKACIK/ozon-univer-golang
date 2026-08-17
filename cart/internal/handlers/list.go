package handlers

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/pkg"
)

type ListHandler struct {
	name string
}

func NewListHandler() *ListHandler {

	return &ListHandler{
		name: "list handler",
	}
}

type ListRequest struct {
	User int64 `json:"user,omitempty"`
}

func (r ListRequest) Validate() error {

	return nil
}

type ListItemResponse struct {
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Price uint32 `json:"price,omitempty"`
}

type ListResponse struct {
	Items      []ListItemResponse `json:"items,omitempty"`
	TotalPrice uint32             `json:"total_price,omitempty"`
}

func (h ListHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := ListRequest{}
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
	for item in items:
		product.get_product(sku) < name,price

	if ok:
		return response
	if fail:
		return error

	*/

	pkg.GetSuccessResponse(w, http.StatusOK)
}
