package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (svc *Service) CancelOrder(ctx context.Context, orderID int64) error {
	err := svc.tx.RunSerializable(ctx, func(ctxTx context.Context) error {
		// set order status as canceled
		err := svc.repo.Order.SetOrderStatus(ctxTx, orderID, OrderStatusCanceled)
		if err != nil {
			return fmt.Errorf("setOrderStatus: %v", err)
		}

		// get reservation
		reservation, err := svc.repo.Reservation.GetReservation(ctxTx, orderID)
		if err != nil {
			return fmt.Errorf("getReservation: %v", err)
		}

		// return reserved products to stocks
		for _, item := range reservation.Stocks {
			err := svc.repo.Stock.ModifyStocksCount(ctxTx, item.Stock, item.SKU, operationAdd)
			if err != nil {
				return fmt.Errorf("modifyStocksCount: %v", err)
			}
		}

		// delete reservation
		err = svc.repo.Reservation.DeleteReservation(ctxTx, orderID)
		if err != nil {
			return fmt.Errorf("deleteReservation: %w", err)
		}
		return nil
	})
	if err != nil {
		return status.Errorf(codes.Internal, "transaction serializable: %v", err)
	}
	return nil
}
