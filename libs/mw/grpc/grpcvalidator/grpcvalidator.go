package grpcvalidator

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator interface {
	ValidateAll() error
}

func Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if validator, ok := req.(Validator); ok {
		if err := validator.ValidateAll(); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
	}
	return handler(ctx, req)
}
