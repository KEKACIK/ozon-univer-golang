package item

import (
	"encoding/json"
	"errors"
	"net/http"
	"route256/cart/internal/pkg"
)

type Adder interface {
	Add(User int64, SKU uint32, Count uint16) error
}

type AddHandler struct {
	name string

	itemAdder Adder
}

func NewAddHandler(itemAdder Adder) *AddHandler {

	return &AddHandler{
		name:      "item add handler",
		itemAdder: itemAdder,
	}
}

type AddRequest struct {
	User  int64
	SKU   uint32
	Count uint16
}

var (
	ErrIncorrectUser  = errors.New("incorrect user")
	ErrIncorrectSKU   = errors.New("incorrect SKU")
	ErrIncorrectCount = errors.New("incorrect quantity")
)

func (r AddRequest) Validate() error {
	if r.User <= 0 {
		return ErrIncorrectUser
	}
	if r.SKU == 0 {
		return ErrIncorrectSKU
	}
	if r.Count == 0 {
		return ErrIncorrectCount
	}

	return nil
}

func (h AddHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := AddRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		pkg.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

}
