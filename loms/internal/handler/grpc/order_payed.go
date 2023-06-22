package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"route256/loms/pkg/loms_v1"
)

func (h *Handler) OrderPayed(ctx context.Context, in *loms_v1.OrderPayedRequest) (*emptypb.Empty, error) {
	err := h.impl.OrderPayed(ctx, in.OrderId)
	return &emptypb.Empty{}, err
}
