package orders

import (
	"encoding/json"
	"net/http"
	"route256/loms/internal/models"
	"route256/loms/internal/pkg"
)

type InfoService interface {
	InfoOrder(orderID int64) (models.OrderModel, error)
}

type InfoHandler struct {
	name        string
	infoService InfoService
}

func NewInfoHandler(infoService InfoService) *InfoHandler {

	return &InfoHandler{
		name:        "orders info handler",
		infoService: infoService,
	}
}

type InfoRequest struct {
	OrderId int64 `json:"order_id,omitempty"`
}

func (r InfoRequest) Validate() error {

	return nil
}

type InfoItemResponse struct {
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

type InfoResponse struct {
	Id     int64              `json:"id,omitempty"`
	Status string             `json:"status,omitempty"`
	User   int64              `json:"user,omitempty"`
	Items  []InfoItemResponse `json:"items,omitempty"`
}

func (r *InfoResponse) Init(order models.OrderModel) {
	r.Id = order.ID
	r.Status = string(order.Status)
	r.User = order.UserID

	for _, v := range order.Items {
		r.Items = append(r.Items, InfoItemResponse{
			SKU:   v.SKU,
			Count: v.Count,
		})
	}
}

func (h InfoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !pkg.CheckMethodPost(w, r) {
		return
	}

	req := &InfoRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	order, err := h.infoService.InfoOrder(req.OrderId)
	if err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	infoResponse := InfoResponse{}
	infoResponse.Init(order)

	raw, err := json.Marshal(infoResponse)
	if err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	pkg.GetSuccessResponseWithBody(w, raw, http.StatusOK)
}
