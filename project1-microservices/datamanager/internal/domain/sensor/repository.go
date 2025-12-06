package sensor

import "context"

type Repository interface {
	GetById(ctx context.Context, id int32) (*Reading, error)
	List(ctx context.Context) ([]Reading, error)
	Create(ctx context.Context, reading *Reading) error
	Update(ctx context.Context, id int32, reading *Reading) error
	Delete(ctx context.Context, id int32) error
}
