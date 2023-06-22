package grpcratelimitter

import (
	"context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRateLimiterUnaryInterceptor(limiter *rate.Limiter) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {

		// Wait for a rate-limited operation to be allowed
		if err := limiter.Wait(ctx); err != nil {
			return status.Errorf(codes.ResourceExhausted, err.Error())
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
