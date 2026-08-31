package loms

import (
	"context"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/models"
	loms_v1 "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/loms/v1"
)

type Client struct {
	loms_v1.LomsServiceClient
}

func NewClient(c loms_v1.LomsServiceClient) *Client {
	return &Client{LomsServiceClient: c}
}

func (c *Client) Create(
	ctx context.Context,
	user int64,
	items []*models.CartModel,
) (int64, error) {
	resp, err := c.LomsServiceClient.OrderCreate(
		ctx,
		&loms_v1.OrderCreateRequest{
			User:  user,
			Items: OrderCreateModel2Request(items),
		},
	)
	if err != nil {
		return 0, err
	}

	return resp.OrderId, nil
}

func (c *Client) GetStocks(ctx context.Context, sku uint32) (uint64, error) {
	resp, err := c.LomsServiceClient.GetStocks(ctx, &loms_v1.GetStocksRequest{Sku: sku})
	if err != nil {
		return 0, err
	}

	return resp.Count, nil
}
