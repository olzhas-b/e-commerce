package grpclogger

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
)

const requestKey = "grpc_request_id"

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, id := contextWithRequestID(ctx)
	log.Printf("[grpc request] method: %s, req: %v requestID: %s, \n", info.FullMethod, req, id)
	response, err := handler(ctx, req)
	log.Printf("[grpc response] method: %s, resp: %v,  requestID: %s, err: %v\n", info.FullMethod, response, id, err)

	return response, err
}

// TODO: check why we can't take requestID from context
// every time we generate a new id
func contextWithRequestID(ctx context.Context) (context.Context, string) {
	id, ok := ctx.Value(requestKey).(string)
	if !ok {
		id = uuid.NewString()
		ctx = context.WithValue(ctx, requestKey, id)
	}
	return ctx, id
}
