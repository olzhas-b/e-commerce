package server

import (
	"route256/checkout/internal/model"
	"route256/checkout/pkg/checkout_v1"
)

func ItemsToReq(items []model.Item) []*checkout_v1.Item {
	response := make([]*checkout_v1.Item, 0, len(items))
	for _, item := range items {
		response = append(response, itemToReq(&item))
	}
	return response
}

func itemToReq(item *model.Item) *checkout_v1.Item {
	return &checkout_v1.Item{
		Sku:   item.SKU,
		Name:  item.Name,
		Count: uint32(item.Count),
		Price: item.Price,
	}
}
