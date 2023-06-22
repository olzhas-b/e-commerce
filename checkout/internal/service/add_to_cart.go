package service

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256/checkout/internal/model"
)

var ErrStockInsufficient = errors.New("stock insufficient")

func (svc *Service) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	if err := svc.checkAmountOfProduct(ctx, sku, count); err != nil {
		return status.Errorf(codes.Internal, "[checkAmountOfProduct] error: %v", err)
	}

	item := model.Item{SKU: sku, Count: count}
	if err := svc.repo.Cart.AddToCart(ctx, user, item); err != nil {
		return status.Errorf(codes.Internal, "[AddToCart] error: %v", err)
	}
	return nil
}

func (svc *Service) checkAmountOfProduct(ctx context.Context, sku uint32, count uint16) error {
	stocks, err := svc.loms.Stocks(ctx, sku)
	if err != nil {
		return fmt.Errorf("stocks: %w", err)
	}

	counter := int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			return nil
		}
	}

	return ErrStockInsufficient
}
