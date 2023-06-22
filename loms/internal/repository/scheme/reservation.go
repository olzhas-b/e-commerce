package scheme

type Reservation struct {
	OrderID int64
	Stocks  []StockReservation
}

type StockReservation struct {
	SKU int64
	Stock
}
