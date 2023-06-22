package inmemory

import (
	"context"
	"route256/checkout/internal/model"
	"route256/libs/postgresdb"
)

type CartRepository struct {
	*postgresdb.Postgres
	inmemory *InMemoryStoraage
}

func (repo *CartRepository) GetItemCount(ctx context.Context, userID int64, sku uint32) (uint16, error) {
	panic("implement me")
}

func NewCartRepository(inmemory *InMemoryStoraage) *CartRepository {
	return &CartRepository{inmemory: inmemory}
}

func (repo *CartRepository) AddToCart(ctx context.Context, userID int64, item model.Item) error {
	return repo.inmemory.Add(ctx, userID, item)
}

func (repo *CartRepository) DeleteFromCart(ctx context.Context, userID int64, item model.Item) error {
	return repo.inmemory.Delete(ctx, userID, item)
}

func (repo *CartRepository) GetCartItems(ctx context.Context, userID int64) ([]model.Item, error) {
	return repo.inmemory.Get(ctx, userID), nil
}

func (repo *CartRepository) UpdateItemFromCart(ctx context.Context, userID int64, status int16, item model.Item) error {
	return repo.inmemory.Delete(ctx, userID, item)
}
