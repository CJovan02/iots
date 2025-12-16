package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryServerLoggingInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {

	start := time.Now()

	log.Printf("[gRPC] START method=%s", info.FullMethod)

	// Call of the RPC handler
	resp, err = handler(ctx, req)

	duration := time.Since(start)

	if err != nil {
		st, _ := status.FromError(err)
		log.Printf(
			"[gRPC] ERROR method=%s, duration=%s, code=%s, msg=%s",
			info.FullMethod,
			duration,
			st.Code().String(),
			st.Message(),
		)
	} else {
		log.Printf("[gRPC] END method=%s duration=%s", info.FullMethod, duration)
	}

	return resp, err
}
