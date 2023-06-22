package service

import (
	"context"
	"route256/checkout/internal/model"
	"route256/checkout/internal/repository"
	"route256/libs/tx"
)

type Service struct {
	loms           LOMS
	productService ProductService

	tx   tx.TransactionManager
	repo *repository.Repository
}

func New(repo *repository.Repository, tx tx.TransactionManager, loms LOMS, productService ProductService) *Service {
	return &Service{
		loms:           loms,
		productService: productService,
		repo:           repo,
		tx:             tx,
	}
}

type LOMS interface {
	CreateOrder(ctx context.Context, userID int64, items []model.Item) (int64, error)
	Stocks(ctx context.Context, sku uint32) ([]model.Stock, error)
}

type ProductService interface {
	GetProduct(ctx context.Context, sku uint32) (model.Product, error)
}
