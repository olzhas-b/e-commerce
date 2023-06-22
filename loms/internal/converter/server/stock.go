package server

import (
	"route256/loms/internal/model"
	"route256/loms/pkg/loms_v1"
)

func StocksToResp(stocks []model.Stock) *loms_v1.StocksResponse {
	var resp loms_v1.StocksResponse
	for _, stock := range stocks {
		resp.Stocks = append(resp.Stocks, StockToResp(stock))
	}
	return &resp
}

func StockToResp(stock model.Stock) *loms_v1.Stock {
	return &loms_v1.Stock{
		Count:       uint64(stock.Count),
		WarehouseId: stock.WarehouseID,
	}
}
