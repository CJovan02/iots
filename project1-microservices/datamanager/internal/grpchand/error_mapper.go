package grpchand

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MapErrToGrpc it appends grpc status code to each error type that it's trying to handle
func MapErrToGrpc(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return status.Errorf(codes.NotFound, err.Error())
	default:
		return status.Errorf(codes.Internal, "internal server error: %v", err)

	}
}
