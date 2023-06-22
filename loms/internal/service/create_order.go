package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"route256/loms/internal/model"
)

const (
	OrderStatusAwaiting = "awaiting"
	OrderStatusFailed   = "failed"
	OrderStatusCanceled = "canceled"
	OrderStatusPayed    = "payed"
)

func (svc *Service) CreateOrder(ctx context.Context, order model.Order) (orderID int64, err error) {
	err = svc.tx.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		orderID, err = svc.repo.Order.CreateOrder(ctxTx, order)
		if err != nil {
			return fmt.Errorf("createOrder: %w", err)
		}

		order.ID = orderID
		err := svc.repo.Order.CreateOrderItems(ctxTx, order)
		if err != nil {
			return fmt.Errorf("createOrderItems: %w", err)
		}
		return nil
	})
	if err != nil {
		return 0, status.Errorf(codes.Internal, "transaction repeatable read: %v", err)
	}

	defer func() {
		if err != nil {
			if err := svc.repo.Order.SetOrderStatus(ctx, orderID, OrderStatusFailed); err != nil {
				log.Printf("setOrderStatus: %v", err)
			}
		}
	}()

	err = svc.tx.RunSerializable(ctx, func(ctxTx context.Context) error {
		for _, item := range order.Items {
			stocks, err := svc.repo.Stock.GetStocks(ctxTx, item.SKU)
			if err != nil {
				return fmt.Errorf("getStocks: %w", err)
			}
			if !svc.checkStockEnough(stocks, int32(item.Count)) {
				return fmt.Errorf("not enough stock")
			}

			err = svc.ReserveOrder(ctxTx, orderID, stocks, item)
			if err != nil {
				return fmt.Errorf("reverseOrder: %w", err)
			}
		}

		err = svc.repo.Order.SetOrderStatus(ctxTx, orderID, OrderStatusAwaiting)
		if err != nil {
			return fmt.Errorf("setOrderStatus: %w", err)
		}

		return nil
	})
	if err != nil {
		return 0, status.Errorf(codes.Internal, "transaction serializable: %v", err)
	}
	return orderID, nil
}

func (svc *Service) checkStockEnough(stocks []model.Stock, need int32) bool {
	var total int32
	for i := 0; i < len(stocks) && total < need; i++ {
		total += stocks[i].Count
	}

	return total >= need
}
