package orders

import (
	"encoding/json"
	"net/http"
	"route256/loms/internal/models"
	"route256/loms/internal/pkg"
)

type CreateService interface {
	CreateOrder(userID int64, items []models.OrderItemModel) (int64, error)
}

type CreateHandler struct {
	name string

	createService CreateService
}

func NewCreateHandler(createService CreateService) *CreateHandler {

	return &CreateHandler{
		name:          "orders create handler",
		createService: createService,
	}
}

type CreateItemRequest struct {
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

type CreateRequest struct {
	User  int64               `json:"user,omitempty"`
	Items []CreateItemRequest `json:"items,omitempty"`
}

func (r CreateRequest) Validate() error {

	return nil
}

func (r CreateRequest) transformToItem() []models.OrderItemModel {
	result := []models.OrderItemModel{}
	for _, v := range r.Items {
		result = append(result, models.OrderItemModel{SKU: v.SKU, Count: v.Count})
	}

	return result
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

	orderID, err := h.createService.CreateOrder(req.User, req.transformToItem())
	if err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	createResponse := CreateResponse{
		OrderId: uint64(orderID),
	}

	raw, err := json.Marshal(createResponse)
	if err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	pkg.GetSuccessResponseWithBody(w, raw, http.StatusOK)
}
