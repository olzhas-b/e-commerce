package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"route256/checkout/internal/model"
	"route256/libs/workerpool"
)

func (svc *Service) ListCart(ctx context.Context, userID int64) (totalPrice uint32, items []model.Item, err error) {
	items, err = svc.repo.Cart.GetCartItems(ctx, userID)
	if err != nil {
		return 0, nil, status.Errorf(codes.Internal, "[GetCartItems] error: %v", err)
	}
	svc.fillItemsField(ctx, items)
	totalPrice = svc.getTotalPrice(ctx, items)
	return totalPrice, items, nil
}

func (svc *Service) fillItemsField(ctx context.Context, items []model.Item) {
	pool := workerpool.NewAsync(ctx, 10, 100)
	for i := 0; i < len(items); i++ {
		i := i
		_ = pool.Exec(ctx, func(ctx context.Context) error {
			product, err := svc.productService.GetProduct(ctx, items[i].SKU)
			if err != nil {
				log.Printf("[fillItemsField] couldn't get product by sku: %d, error: %v\n", items[i].SKU, err)
			}
			items[i].Price = product.Price
			items[i].Name = product.Name
			return nil
		})
	}
	pool.Wait()
}

func (svc *Service) getTotalPrice(ctx context.Context, items []model.Item) uint32 {
	var totalPrice uint32
	for _, item := range items {
		totalPrice += item.Price * uint32(item.Count)
	}
	return totalPrice
}
