package service

import (
	"route256/libs/tx"
	"route256/loms/internal/repository"
)

type Service struct {
	tx   *tx.Manager
	repo *repository.Repository
}

func New(repo *repository.Repository, tx *tx.Manager) *Service {
	return &Service{
		repo: repo,
		tx:   tx,
	}
}
