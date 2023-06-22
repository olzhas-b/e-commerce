package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256/loms/internal/model"
)

func (svc *Service) ListOrder(ctx context.Context, orderID int64) (model.Order, error) {
	order, err := svc.repo.Order.GetOrder(ctx, orderID)
	if err != nil {
		return model.Order{}, status.Errorf(codes.Internal, "getOrder: %v", err)
	}
	return order, nil
}
