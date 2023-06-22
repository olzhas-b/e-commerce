package server

import (
	"route256/checkout/internal/model"
	"route256/checkout/pkg/checkout_v1"
)

func ListCartResponse(totalPrice uint32, items []model.Item) *checkout_v1.ListCartResponse {
	return &checkout_v1.ListCartResponse{
		TotalPrice: totalPrice,
		Items:      ItemsToReq(items),
	}
}
