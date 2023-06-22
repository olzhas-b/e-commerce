package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"route256/loms/pkg/loms_v1"
)

func (h *Handler) CancelOrder(ctx context.Context, in *loms_v1.CancelOrderRequest) (*emptypb.Empty, error) {
	err := h.impl.CancelOrder(ctx, in.OrderId)
	return &emptypb.Empty{}, err
}
