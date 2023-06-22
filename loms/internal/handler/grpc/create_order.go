package grpc

import (
	"context"
	"route256/loms/internal/converter/server"
	"route256/loms/pkg/loms_v1"
)

func (h *Handler) CreateOrder(ctx context.Context, in *loms_v1.CreateOrderRequest) (*loms_v1.CreateOrderResponse, error) {
	order := server.OrderFromReq(in)
	orderID, err := h.impl.CreateOrder(ctx, order)
	return server.OrderIDToResp(orderID), err
}
