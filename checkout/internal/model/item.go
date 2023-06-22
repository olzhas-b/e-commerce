package model

type Item struct {
	SKU   uint32
	Count uint16
	Name  string
	Price uint32
}

type ListItems struct {
	Items      []Item
	TotalPrice uint32
}
