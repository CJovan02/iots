package interceptor

import (
	"context"
	"errors"

	"github.com/CJovan02/iots/datamanager/internal/domain/sensor"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryServerErrMappingInterceptor(
	ctx context.Context,
	req any,
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	resp, err = handler(ctx, req)

	if err == nil {
		return resp, nil
	}

	if _, ok := status.FromError(err); ok {
		return resp, err
	}

	return resp, MapErrToGrpc(err)
}

// MapErrToGrpc it appends grpc status code to each error type that it's trying to handle
func MapErrToGrpc(err error) error {
	if err == nil {
		return nil
	}

	var notFound *sensor.NotFound
	var invalid *sensor.InvalidArgument

	switch {
	case errors.Is(err, pgx.ErrNoRows):
	case errors.As(err, &notFound):
		return status.Errorf(codes.NotFound, err.Error())
	case errors.As(err, &invalid):
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	return status.Errorf(codes.Internal, "internal server error: %v", err)
}
