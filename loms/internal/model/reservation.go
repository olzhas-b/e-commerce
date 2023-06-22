package model

type Reservation struct {
	OrderID int64
	Stocks  []StockReservation
}

type StockReservation struct {
	SKU uint32
	Stock
}
