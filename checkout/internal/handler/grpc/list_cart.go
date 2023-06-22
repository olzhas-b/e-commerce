package grpc

import (
	"context"
	"route256/checkout/internal/converter/server"
	"route256/checkout/pkg/checkout_v1"
)

func (svc *Handler) ListCart(ctx context.Context, in *checkout_v1.ListCartRequest) (*checkout_v1.ListCartResponse, error) {
	totalPrice, items, err := svc.impl.ListCart(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return server.ListCartResponse(totalPrice, items), nil
}
