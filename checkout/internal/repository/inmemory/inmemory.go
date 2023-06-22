package inmemory

import (
	"context"
	"errors"
	"route256/checkout/internal/model"
	"sync"
)

type InMemoryStoraage struct {
	storage map[int64]map[uint32]*model.Item
	mu      sync.Mutex
}

func New() *InMemoryStoraage {
	return &InMemoryStoraage{
		storage: make(map[int64]map[uint32]*model.Item),
	}
}

var ErrorNotFound = errors.New("not found")

func (memo *InMemoryStoraage) Add(_ context.Context, userID int64, item model.Item) error {
	memo.mu.Lock()
	defer memo.mu.Unlock()

	if memo.storage[userID] == nil {
		memo.storage[userID] = make(map[uint32]*model.Item)
	}

	if _, ok := memo.storage[userID][item.SKU]; ok {
		memo.storage[userID][item.SKU].Count += item.Count
	} else {
		memo.storage[userID][item.SKU] = &item
	}
	return nil
}

func (memo *InMemoryStoraage) Get(_ context.Context, userID int64) []model.Item {
	memo.mu.Lock()
	defer memo.mu.Unlock()

	var items []model.Item
	for _, item := range memo.storage[userID] {
		items = append(items, *item)
	}
	return items
}

func (memo *InMemoryStoraage) Delete(_ context.Context, userID int64, item model.Item) error {
	memo.mu.Lock()
	defer memo.mu.Unlock()

	m, ok := memo.storage[userID]
	if !ok {
		return ErrorNotFound
	}

	storedItem, ok := m[item.SKU]
	if !ok {
		return ErrorNotFound
	}

	if storedItem.Count <= item.Count {
		delete(m, item.SKU)
	} else {
		m[item.SKU].Count -= item.Count
	}
	return nil
}
