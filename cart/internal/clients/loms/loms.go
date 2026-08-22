package loms

import (
	"context"

	loms_v1 "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/loms/v1"
)

type Client struct {
	loms_v1.LomsClient
}

func NewClient(c loms_v1.LomsClient) *Client {
	return &Client{LomsClient: c}
}

func (c *Client) GetStocks(ctx context.Context, sku uint32) (uint64, error) {
	resp, err := c.LomsClient.GetStocks(ctx, &loms_v1.GetStocksRequest{Sku: sku})
	if err != nil {
		return 0, err
	}

	return resp.Count, nil
}
