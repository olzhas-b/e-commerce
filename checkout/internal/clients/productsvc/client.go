package productsvc

import (
	"context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"route256/checkout/internal/converter/client/productsvc"
	"route256/checkout/internal/model"
	"route256/checkout/pkg/productsvc_v1"
	"route256/libs/mw/grpc/grpcratelimitter"
	"time"
)

type Client struct {
	token string
	impl  productsvc_v1.ProductServiceClient
}

func NewClient(addr, token string, rps int64) *Client {
	limiter := rate.NewLimiter(rate.Every(time.Second/time.Duration(rps)), 1)

	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcratelimitter.NewRateLimiterUnaryInterceptor(limiter)),
	)
	if err != nil {
		panic(err)
	}

	c := productsvc_v1.NewProductServiceClient(conn)
	if err != nil {
		return nil
	}
	return &Client{
		token: token,
		impl:  c,
	}
}

func (c *Client) GetProduct(ctx context.Context, sku uint32) (model.Product, error) {
	response, err := c.impl.GetProduct(ctx, productsvc.GetProductRequest(sku, c.token))
	if err != nil {
		return model.Product{}, err
	}
	return productsvc.ProductFromResp(response), nil
}
