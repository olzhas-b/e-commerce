package loms

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"route256/checkout/internal/converter/client/loms"
	"route256/checkout/internal/model"
	"route256/checkout/pkg/loms_v1"
	"time"
)

type Client struct {
	impl loms_v1.LOMSClient
}

func NewClient(addr string) *Client {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	c := loms_v1.NewLOMSClient(conn)
	return &Client{
		impl: c,
	}
}

func (c *Client) CreateOrder(ctx context.Context, userID int64, items []model.Item) (int64, error) {
	ctx, _ = context.WithTimeout(ctx, time.Second)
	time.Sleep(4 * time.Second)
	resp, err := c.impl.CreateOrder(ctx, loms.CreateOrderRequest(userID, items))
	if err != nil {
		return 0, status.Errorf(codes.Internal, err.Error())
	}
	return loms.OrderIDFromResp(resp), nil
}

func (c *Client) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	resp, err := c.impl.Stocks(ctx, loms.StocksRequest(sku))
	if err != nil {
		return nil, err
	}
	return loms.StocksFromResp(resp), nil
}
