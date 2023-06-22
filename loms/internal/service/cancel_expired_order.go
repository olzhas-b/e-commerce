package service

import (
	"context"
	"fmt"
	"route256/loms/internal/config"
	"time"
)

func (svc *Service) CancelExpiredOrders(ctx context.Context) error {
	expiresTime := time.Millisecond * time.Duration(config.AppConfig.Jobs.CancelOrder.ExpiresTime)
	comparisonTime := time.Now().Add(-expiresTime)
	orderIDs, err := svc.repo.Order.GetExpiredOrderIDs(ctx, comparisonTime, OrderStatusAwaiting)
	if err != nil {
		return fmt.Errorf("getExpiredOrderIDs: %w", err)
	}

	for _, id := range orderIDs {
		if err := svc.CancelOrder(ctx, id); err != nil {
			return fmt.Errorf("cancelOrder: %w", err)
		}
	}
	return nil
}
