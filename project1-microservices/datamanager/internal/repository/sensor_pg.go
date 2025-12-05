package repository

import (
	"context"

	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/domain/sensor"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgRepository struct {
	db *pgxpool.Pool
}

func NewPgRepository(db *pgxpool.Pool) *PgRepository {
	return &PgRepository{db: db}
}

func (r *PgRepository) GetById(ctx context.Context, id int32) (*sensor.SensorReading, error) {
	const query = `
		SELECT *
		FROM sensor_readings
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, id)

	return scanSensorReading(row)
}
