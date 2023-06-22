package scheme

import (
	"route256/loms/internal/model"
	"time"
)

func StockFromScheme(stock Stock) model.Stock {
	return model.Stock{
		WarehouseID: stock.WarehouseID,
		Count:       stock.Count,
	}
}

func StockToScheme(stock model.Stock) Stock {
	return Stock{
		WarehouseID:   stock.WarehouseID,
		Count:         stock.Count,
		LastUpdatedAt: time.Now(),
	}
}

func StocksFromScheme(stocks []Stock) []model.Stock {
	var result []model.Stock
	for _, stock := range stocks {
		result = append(result, StockFromScheme(stock))
	}
	return result
}

func StocksToScheme(stocks []model.Stock) []Stock {
	var result []Stock
	for _, stock := range stocks {
		result = append(result, StockToScheme(stock))
	}
	return result
}

func OrderToScheme(order model.Order) Order {
	var items []Item
	for _, item := range order.Items {
		items = append(items, ItemToScheme(item))
	}
	return Order{
		ID:            OrderIDToScheme(order.ID),
		Price:         order.Price,
		UserID:        order.User,
		Status:        order.Status,
		Items:         items,
		LastUpdatedAt: time.Now(),
	}
}

func OrderFromScheme(order Order) model.Order {
	return model.Order{
		ID:     int64(order.ID),
		User:   order.UserID,
		Price:  order.Price,
		Status: order.Status,
	}
}

func OrderIDsFromScheme(ids []OrderID) []int64 {
	var result []int64
	for _, id := range ids {
		result = append(result, OrderIDFromScheme(id))
	}
	return result
}

func OrderIDFromScheme(id OrderID) int64 {
	return int64(id)
}

func OrderIDToScheme(id int64) OrderID {
	return OrderID(id)
}

func ItemFromScheme(item Item) model.Item {
	return model.Item{
		SKU:   SkuFromScheme(item.SKU),
		Count: uint16(item.Count),
	}
}

func ItemToScheme(item model.Item) Item {
	return Item{
		SKU:   SkuToScheme(item.SKU),
		Count: int32(item.Count),
	}
}

func SkuToScheme(sku uint32) int64 {
	return int64(sku)
}

func SkuFromScheme(sku int64) uint32 {
	return uint32(sku)
}

func ReservationFromScheme(reservation Reservation) model.Reservation {
	var stocks []model.StockReservation
	for _, stock := range reservation.Stocks {
		stocks = append(stocks, StockReservationFromScheme(stock))
	}
	return model.Reservation{
		OrderID: reservation.OrderID,
		Stocks:  stocks,
	}
}

func StockReservationFromScheme(stock StockReservation) model.StockReservation {
	return model.StockReservation{
		SKU:   SkuFromScheme(stock.SKU),
		Stock: StockFromScheme(stock.Stock),
	}
}

func StockReservationToScheme(stock model.StockReservation) StockReservation {
	return StockReservation{
		SKU:   SkuToScheme(stock.SKU),
		Stock: StockToScheme(stock.Stock),
	}
}
