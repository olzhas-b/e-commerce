package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (svc *Service) OrderPayed(ctx context.Context, orderID int64) error {
	err := svc.tx.RunSerializable(ctx, func(ctxTx context.Context) error {
		err := svc.repo.Order.SetOrderStatus(ctxTx, orderID, OrderStatusPayed)
		if err != nil {
			return fmt.Errorf("setOrderStatus: %w", err)
		}

		err = svc.repo.Reservation.DeleteReservation(ctxTx, orderID)
		if err != nil {
			return fmt.Errorf("deleteReservation: %w", err)
		}
		return nil
	})
	if err != nil {
		return status.Errorf(codes.Internal, "transcation repeatable read: %v", err)
	}
	return nil
}
