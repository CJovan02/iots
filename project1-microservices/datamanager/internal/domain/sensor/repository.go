package sensor

import (
	"context"
	"time"
)

type Repository interface {
	GetById(ctx context.Context, id int32) (*Reading, error)
	List(ctx context.Context) ([]Reading, error)
	Create(ctx context.Context, reading *Reading) error
	GetStatistics(ctx context.Context, startTime time.Time, endTime time.Time) (*Statistics, error)
	Update(ctx context.Context, id int32, reading *Reading) error
	Delete(ctx context.Context, id int32) error
}
