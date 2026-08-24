package converter

import (
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/models"
	desc "github.com/KEKACIK/ozon-univer-golang/loms/pkg/api/loms/v1"
)

func OrderConvertCreateRequest2Model(r *desc.OrderCreateRequest) []*models.OrderItemModel {
	result := []*models.OrderItemModel{}
	for _, v := range r.Items {
		result = append(result, &models.OrderItemModel{SKU: v.Sku, Count: uint16(v.Count)})
	}

	return result
}

func OrderConvertModel2InfoResponse(r *models.OrderModel) *desc.OrderInfoResponse {
	result := desc.OrderInfoResponse{
		OrderId: r.ID,
		Status:  string(r.Status),
		User:    r.UserID,
	}
	for _, v := range r.Items {
		result.Items = append(result.Items, &desc.OrderInfoItem{Sku: v.SKU, Count: uint32(v.Count)})
	}

	return &result
}
