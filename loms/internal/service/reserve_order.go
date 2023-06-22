package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256/loms/internal/model"
)

const (
	operationAdd = "+"
	operationSub = "-"
)

func (svc *Service) ReserveOrder(ctx context.Context, orderID int64, stocks []model.Stock, item model.Item) error {
	// get necessary items from stocks to reserve
	stocksForReservation := svc.listOfStockForReservation(ctx, stocks, item)

	// change items count in stocks
	for _, stock := range stocksForReservation {
		err := svc.repo.Stock.ModifyStocksCount(ctx, stock, item.SKU, operationSub)
		if err != nil {
			return status.Errorf(codes.Internal, "modifyStocksCount: %v", err)
		}
	}

	// build reservation
	reservation := model.Reservation{
		OrderID: orderID,
	}
	for _, stock := range stocksForReservation {
		reservation.Stocks = append(reservation.Stocks, model.StockReservation{
			SKU: item.SKU,
			Stock: model.Stock{
				Count:       stock.Count,
				WarehouseID: stock.WarehouseID,
			},
		})
	}

	// create reservation
	err := svc.repo.Reservation.CreateReservation(ctx, reservation)
	if err != nil {
		return status.Errorf(codes.Internal, "createReservation: %v", err)
	}
	return nil
}

func (svc *Service) listOfStockForReservation(ctx context.Context, stocks []model.Stock, item model.Item) []model.Stock {
	response := make([]model.Stock, 0, 1)

	var count = int32(item.Count)
	for i := 0; i < len(stocks) && count > 0; i++ {
		stock := stocks[i]

		if stock.Count >= count {
			stock.Count = count
		}
		count -= stock.Count
		response = append(response, stock)
	}
	return response
}

func (svc *Service) RevertReservation(ctx context.Context, orderID int64) error {
	err := svc.tx.RunSerializable(ctx, func(ctxTx context.Context) error {
		if err := svc.repo.Reservation.DeleteReservation(ctx, orderID); err != nil {
			return fmt.Errorf("deleteReservation: %w", err)
		}
		return nil
	})
	if err != nil {
		return status.Errorf(codes.Internal, "transaction serializable: %v", err)
	}
	return nil
}
