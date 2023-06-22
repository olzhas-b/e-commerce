package grpc

import (
	"route256/loms/internal/service"
	"route256/loms/pkg/loms_v1"
)

type Handler struct {
	impl *service.Service
	loms_v1.UnimplementedLOMSServer
}

func NewHandler(impl *service.Service) *Handler {
	return &Handler{impl: impl}
}
