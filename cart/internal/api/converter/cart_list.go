package converter

import (
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/models"
	desc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/cart/v1"
)

func CartListConvertModel2Response(items []*models.CartModel, totalPrice uint32) *desc.ListResponse {
	result := desc.ListResponse{
		Items:      []*desc.ListItem{},
		TotalPrice: totalPrice,
	}
	for _, item := range items {
		result.Items = append(result.Items, &desc.ListItem{
			Sku:   item.SKU,
			Count: uint32(item.Count),
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return &result
}
