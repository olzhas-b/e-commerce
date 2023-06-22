package server

import "route256/checkout/pkg/checkout_v1"

func OrderIDToPurchaseResponse(orderID int64) *checkout_v1.PurchaseResponse {
	return &checkout_v1.PurchaseResponse{
		OrderId: orderID,
	}
}
