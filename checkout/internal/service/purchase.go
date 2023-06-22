package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	statusPending = iota
	statusSuccess
	statusFailed
)

func (svc *Service) Purchase(ctx context.Context, userID int64) (orderID int64, err error) {
	items, err := svc.repo.Cart.GetCartItems(ctx, userID)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "[Purchase] Couldn't get listCart: %v", err)
	}

	orderID, err = svc.loms.CreateOrder(ctx, userID, items)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "[Purchase] Couldn't createOrder: %v", err)
	}

	err = svc.repo.Cart.DeleteAllFromCart(ctx, userID)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "[Purchase] Couldn't deleteAllFromCart: %v", err)
	}

	return orderID, nil
}
