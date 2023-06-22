package grpc

import (
	"context"
	"route256/loms/internal/converter/server"
	"route256/loms/pkg/loms_v1"
)

func (h *Handler) ListOrder(ctx context.Context, in *loms_v1.ListOrderRequest) (*loms_v1.ListOrderResponse, error) {
	order, err := h.impl.ListOrder(ctx, in.OrderId)
	return server.OrderToListResp(order), err
}
