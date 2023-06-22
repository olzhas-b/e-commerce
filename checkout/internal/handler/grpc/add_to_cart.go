package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"route256/checkout/pkg/checkout_v1"
)

func (svc *Handler) AddToCart(ctx context.Context, in *checkout_v1.AddToCartRequest) (*emptypb.Empty, error) {
	err := svc.impl.AddToCart(ctx, in.UserId, in.Sku, uint16(in.Count))
	return &emptypb.Empty{}, err
}
