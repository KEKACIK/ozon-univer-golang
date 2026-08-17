package loms

import (
	"fmt"
	"net/url"
)

type Client struct {
	name string
	path string
}

func New(name, basePath string) (*Client, error) {
	const handlerName = "get_product"
	path, err := url.JoinPath(basePath, handlerName)
	if err != nil {
		return nil, fmt.Errorf("%s: incorrect base path: %w", name, err)
	}

	return &Client{
		name: name,
		path: path,
	}, nil
}

func (c Client) GetStocks(sku uint32) (uint64, error) {
	return 0, nil
}
