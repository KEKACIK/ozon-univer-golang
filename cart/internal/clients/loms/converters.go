package loms

import (
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/models"
	loms_v1 "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/loms/v1"
)

func OrderCreateModel2Request(items []*models.CartModel) []*loms_v1.OrderCreateItem {
	result := []*loms_v1.OrderCreateItem{}

	for _, item := range items {
		result = append(result, &loms_v1.OrderCreateItem{
			Sku:   item.SKU,
			Count: uint32(item.Count),
		})
	}

	return result
}
