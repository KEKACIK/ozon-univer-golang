// ProductClient к сожалению не работает на 08.2026
// Поэтому всё захардкожено (харкодено хз как правильно)
package products

import (
	"context"
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

func (c Client) GetProductInfo(ctx context.Context, sku uint32) (string, uint32, error) {
	switch sku {
	case uint32(1076963):
		return "Теория нравственных чувств | Смит Адам", 1_150, nil
	case uint32(5678901):
		return "Орден кольца | Толкин", 8_900, nil
	case uint32(4200000):
		return "Микрофон Razer | USB", 5_399, nil
	case uint32(2500001):
		return "Iphone 15 Pro Max | Паленый", 32_700, nil
	case uint32(8765432):
		return "Redmi 8X Pro Gaming Max Plus 32/1024", 32_659, nil
	case uint32(7000123):
		return "Кукла вуду для Артура", 501, nil
	}

	return "", 0, ErrNotFound
}

func (c Client) GetListSKUs(ctx context.Context) ([]uint32, error) {

	return []uint32{
		1076963,
		5678901,
		4200000,
		2500001,
		8765432,
		7000123,
	}, nil
}
