package sensor

import (
	"context"
	"time"
)

type Repository interface {
	CountAll(ctx context.Context) (*uint32, error)
	Exists(ctx context.Context, id uint32) (bool, error)
	GetById(ctx context.Context, id uint32) (*Reading, error)
	List(ctx context.Context, offset uint32, limit uint32) ([]Reading, error)
	GetStatistics(ctx context.Context, startTime time.Time, endTime time.Time) (*Statistics, error)
	Create(ctx context.Context, reading *Reading) (*uint32, error)
	BatchCreate(ctx context.Context, readings []*Reading) ([]uint32, error)
	Update(ctx context.Context, id uint32, reading *Reading) error
	Delete(ctx context.Context, id uint32) error
}
