package orders

import (
	"encoding/json"
	"net/http"
	"route256/loms/internal/pkg"
)

type InfoHandler struct {
	name string
}

func NewInfoHandler() *InfoHandler {

	return &InfoHandler{
		name: "orders info handler",
	}
}

type InfoRequest struct {
	OrderId int64
}

func (r InfoRequest) Validate() error {

	return nil
}

type InfoItemResponse struct {
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

type InfoResponse struct {
	Status string             `json:"status,omitempty"`
	User   int64              `json:"user,omitempty"`
	Items  []InfoItemResponse `json:"items,omitempty"`
}

func (h InfoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &InfoRequest{}
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

	if ok:
		return response
	if failt:
		return error (not found)

	*/

	infoResponse := InfoResponse{
		Status: "status",
		User:   1,
		Items: []InfoItemResponse{
			{
				SKU:   1,
				Count: 1,
			},
		},
	}

	raw, err := json.Marshal(infoResponse)
	if err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	pkg.GetSuccessResponseWithBody(w, raw, http.StatusOK)
}
