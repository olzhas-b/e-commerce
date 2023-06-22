package server

import (
	"route256/loms/internal/model"
	"route256/loms/pkg/loms_v1"
)

func ItemsToReq(items []model.Item) []*loms_v1.Item {
	response := make([]*loms_v1.Item, 0, len(items))
	for _, item := range items {
		response = append(response, itemToReq(&item))
	}
	return response
}

func itemToReq(item *model.Item) *loms_v1.Item {
	return &loms_v1.Item{
		Sku:   item.SKU,
		Count: uint32(item.Count),
	}
}

func itemFromReq(item *loms_v1.Item) model.Item {
	return model.Item{
		SKU:   item.Sku,
		Count: uint16(item.Count),
	}
}
