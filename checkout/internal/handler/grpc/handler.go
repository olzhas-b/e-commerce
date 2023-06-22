package grpc

import (
	"route256/checkout/internal/service"
	"route256/checkout/pkg/checkout_v1"
)

type Handler struct {
	impl *service.Service
	checkout_v1.UnimplementedCheckoutServer
}

func NewHandler(impl *service.Service) *Handler {
	return &Handler{impl: impl}
}
