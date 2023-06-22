package model

type Stock struct {
	WarehouseID int64
	Count       uint64
}

type StockRequest struct {
	SKU uint32
}
