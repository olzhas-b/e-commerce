package loms

import (
	"route256/checkout/internal/model"
	"route256/checkout/pkg/loms_v1"
)

func CreateOrderRequest(userID int64, items []model.Item) *loms_v1.CreateOrderRequest {
	return &loms_v1.CreateOrderRequest{
		Order: &loms_v1.Order{
			UserId: userID,
			Items:  ItemsToReq(items),
		},
	}
}

func OrderIDFromResp(resp *loms_v1.CreateOrderResponse) int64 {
	return resp.OrderId
}

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

func StocksRequest(sku uint32) *loms_v1.StocksRequest {
	return &loms_v1.StocksRequest{
		Sku: sku,
	}
}

func StocksFromResp(resp *loms_v1.StocksResponse) []model.Stock {
	stocks := make([]model.Stock, 0, len(resp.Stocks))
	for _, rs := range resp.Stocks {
		stocks = append(stocks, StockFromResp(rs))
	}
	return stocks
}

func StockFromResp(stock *loms_v1.Stock) model.Stock {
	return model.Stock{
		WarehouseID: stock.WarehouseId,
		Count:       stock.Count,
	}
}
