package repository

import (
	"context"

	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/domain/sensor"
	"github.com/jackc/pgx/v5"
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

func (r *PgRepository) List(ctx context.Context) ([]sensor.SensorReading, error) {
	const query = `
		SELECT *
		FROM sensor_readings
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	readings, err := pgx.CollectRows(rows, scanSensorReadingValue)
	if err != nil {
		return nil, err
	}

	return readings, nil
}

func (r *PgRepository) Create(ctx context.Context, reading *sensor.SensorReading) error {
	const query = `
		INSERT INTO sensor_readings 
		(timestamp, temperature, humidity, tvoc, e_co2, raw_hw, raw_ethanol, pm_25, fire_alarm)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(ctx, query, sensorReadingArgs(reading)...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PgRepository) Update(ctx context.Context, id int32, reading *sensor.SensorReading) error {
	const query = `
		UPDATE sensor_readings
		SET 
		    timestamp = $2,
		    temperature = $3,
		    humidity = $4,
		    tvoc = $5,
		    e_co2 = $6,
		    raw_hw = $7,
		    raw_ethanol = $8,
		    pm_25 = $9,
		    fire_alarm = $10
		WHERE id = $1
	`
	args := append([]any{id}, sensorReadingArgs(reading)...)

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PgRepository) Delete(ctx context.Context, id int32) error {
	const query = `
		DELETE FROM sensor_readings
		WHERE id = $1
	`

	cmd, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
