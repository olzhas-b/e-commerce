package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256/loms/internal/model"
)

func (svc *Service) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	stocks, err := svc.repo.Stock.GetStocks(ctx, sku)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getStocks: %v", err)
	}
	return stocks, nil
}
