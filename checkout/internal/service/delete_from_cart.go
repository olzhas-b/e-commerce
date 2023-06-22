package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256/checkout/internal/model"
)

func (svc *Service) DeleteFromCart(ctx context.Context, userID int64, sku uint32, count uint16) error {
	itemCount, err := svc.repo.Cart.GetItemCount(ctx, userID, sku)
	if err != nil {
		return status.Errorf(codes.Internal, "[GetItemCount] error: %v", err)
	}

	item := model.Item{SKU: sku, Count: itemCount - count}
	if count >= itemCount {
		if err := svc.repo.Cart.DeleteFromCart(ctx, userID, item.SKU); err != nil {
			return status.Errorf(codes.Internal, "[DeleteFromCart] error: %v", err)
		}
	} else {
		if err := svc.repo.Cart.UpdateItemFromCart(ctx, userID, item); err != nil {
			return status.Errorf(codes.Internal, "[UpdateItemFromCart] error: %v", err)
		}
	}

	return nil
}
