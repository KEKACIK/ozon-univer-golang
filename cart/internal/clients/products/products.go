// ProductClient к сожалению не работает на 08.2026
// Поэтому всё захардкожено (харкодено хз как правильно)
package products

import (
	"errors"
)

type Client struct {
	name string
	path string
}

func New(name, basePath string) (*Client, error) {
	return &Client{
		name: name,
		path: basePath,
	}, nil
}

var ErrNotFound = errors.New("not found")

func (c Client) GetProductInfo(sku uint32) (string, uint32, error) {
	switch sku {
	case uint32(1076963):
		return "Теория нравственных чувств | Смит Адам", 1150, nil
	}

	return "", 0, ErrNotFound
}

func (c Client) GetListSKUs() ([]uint32, error) {

	return []uint32{
		1076963,
	}, nil
}
