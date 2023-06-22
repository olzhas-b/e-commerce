package repository

import (
	"context"
	"route256/libs/postgresdb"
	"route256/loms/internal/model"
	"route256/loms/internal/repository/postgres/order"
	"route256/loms/internal/repository/postgres/reservation"
	"route256/loms/internal/repository/postgres/stock"
	"time"
)

type Repository struct {
	Stock       Stock
	Order       Order
	Reservation Reservation
}

func NewRepository(db *postgresdb.Postgres) *Repository {
	return &Repository{
		Stock:       stock.NewStocksRepository(db),
		Order:       order.NewOrderRepository(db),
		Reservation: reservation.NewReservationRepository(db),
	}
}

type Stock interface {
	GetStocks(ctx context.Context, sku uint32) ([]model.Stock, error)
	ModifyStocksCount(ctx context.Context, stock model.Stock, sku uint32, operation string) error
}

type Reservation interface {
	GetReservation(ctx context.Context, orderID int64) (model.Reservation, error)
	CreateReservation(ctx context.Context, reservation model.Reservation) error
	DeleteReservation(ctx context.Context, orderID int64) error
}

type Order interface {
	GetOrder(ctx context.Context, orderID int64) (model.Order, error)
	GetExpiredOrderIDs(ctx context.Context, comparisonTime time.Time, status string) ([]int64, error)
	CreateOrder(ctx context.Context, order model.Order) (orderID int64, err error)
	SetOrderStatus(ctx context.Context, orderID int64, status string) error

	CreateOrderItems(ctx context.Context, order model.Order) error
}
