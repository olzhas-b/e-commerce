package grpc

import (
	"context"
	"route256/checkout/internal/converter/server"
	"route256/checkout/pkg/checkout_v1"
)

func (svc *Handler) Purchase(ctx context.Context, in *checkout_v1.PurchaseRequest) (*checkout_v1.PurchaseResponse, error) {
	orderID, err := svc.impl.Purchase(ctx, in.UserId)
	return server.OrderIDToPurchaseResponse(orderID), err
}
