package grpc

import (
	"context"
	"route256/loms/internal/converter/server"
	"route256/loms/pkg/loms_v1"
)

func (h *Handler) Stocks(ctx context.Context, in *loms_v1.StocksRequest) (*loms_v1.StocksResponse, error) {
	stocks, err := h.impl.Stocks(ctx, in.Sku)
	return server.StocksToResp(stocks), err
}
