package scheme

import "route256/checkout/internal/model"

func ItemFromScheme(item Item) model.Item {
	return model.Item{
		SKU:   uint32(item.SKU),
		Count: uint16(item.Count),
		Price: uint32(item.Price),
	}
}

func ItemToScheme(item model.Item) Item {
	return Item{
		SKU:   int64(item.SKU),
		Count: int32(item.Count),
		Price: int64(item.Price),
	}
}

func ItemsFromScheme(items []Item) []model.Item {
	var result []model.Item
	for _, item := range items {
		result = append(result, ItemFromScheme(item))
	}
	return result
}

func ItemsToScheme(items []model.Item) []Item {
	var result []Item
	for _, item := range items {
		result = append(result, ItemToScheme(item))
	}
	return result
}

func SkuToScheme(sku uint32) int64 {
	return int64(sku)
}

func SkuFromScheme(sku int64) uint32 {
	return uint32(sku)
}
