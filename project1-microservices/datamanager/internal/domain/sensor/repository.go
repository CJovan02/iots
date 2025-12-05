package sensor

import "context"

type Repository interface {
	GetById(ctx context.Context, id int32) (*SensorReading, error)
	List(ctx context.Context) ([]SensorReading, error)
	Create(ctx context.Context, reading *SensorReading) error
	Update(ctx context.Context, id int32, reading *SensorReading) error
	Delete(ctx context.Context, id int32) error
}
