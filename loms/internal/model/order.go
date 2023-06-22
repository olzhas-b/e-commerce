package model

type Order struct {
	ID     int64
	User   int64
	Price  int64
	Items  []Item
	Status string
}
