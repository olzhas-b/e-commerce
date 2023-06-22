package repository

import (
	"context"
	"route256/checkout/internal/model"
	"route256/checkout/internal/repository/postgres/cart"
	"route256/checkout/internal/repository/postgres/purchase"
	"route256/libs/postgresdb"
	"route256/libs/tx"
)

type Repository struct {
	Cart     Cart
	Purchase Purchase
	tx       tx.TransactionManager
}

func NewRepository(db *postgresdb.Postgres) *Repository {
	return &Repository{
		Cart:     cart.NewCartRepository(db),
		Purchase: purchase.NewPurchaseRepository(db),
	}
}

type Cart interface {
	GetItemCount(ctx context.Context, userID int64, sku uint32) (uint16, error)
	GetCartItems(ctx context.Context, userID int64) ([]model.Item, error)
	AddToCart(ctx context.Context, userID int64, item model.Item) error
	UpdateItemFromCart(ctx context.Context, userID int64, item model.Item) error
	DeleteFromCart(ctx context.Context, userID int64, sku uint32) error
	DeleteAllFromCart(ctx context.Context, userID int64) error
}

type Purchase interface {
	Create(ctx context.Context, userID, orderID int64, status int16) (purchaseID int64, err error)
	UpdateStatus(ctx context.Context, userID, orderID int64, status int16) error
	AddToPurchaseItems(ctx context.Context, purchaseID int64, items []model.Item) error
}
