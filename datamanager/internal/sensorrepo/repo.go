package sensorrepo

import (
	"context"
	"time"

	"github.com/CJovan02/iots/datamanager/internal/domain/sensor"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CountAll(ctx context.Context) (*uint32, error) {
	const query = `SELECT COUNT(*) FROM sensor_readings`

	var count uint32
	err := r.db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func (r *Repository) List(ctx context.Context, offset uint32, limit uint32) ([]sensor.Reading, error) {
	const query = `
		SELECT *
		FROM sensor_readings
		ORDER BY timestamp
		OFFSET $1
		LIMIT $2
	`

	rows, err := r.db.Query(ctx, query, offset, limit)
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

func (r *Repository) Exists(ctx context.Context, id uint32) (bool, error) {
	const query = `
		SELECT EXISTS (
			SELECT 1
			FROM sensor_readings
			WHERE id = $1   
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *Repository) GetById(ctx context.Context, id uint32) (*sensor.Reading, error) {
	const query = `
		SELECT *
		FROM sensor_readings
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, id)

	return scanSensorReading(row)
}

func (r *Repository) GetStatistics(ctx context.Context, startTime time.Time, endTime time.Time) (*sensor.Statistics, error) {
	const query = `
		SELECT 
			COUNT(*) as readings_count, 
			COALESCE(MIN(temperature), 0) as min_temp,
			COALESCE(MAX(temperature), 0) as max_temp,
			COALESCE(AVG(temperature), 0) as avg_temp,
			COALESCE(MIN(humidity), 0) as min_humidity,
			COALESCE(MAX(humidity), 0) as max_humidity,
			COALESCE(AVG(humidity), 0) as avg_humidity,
			COALESCE(SUM(tvoc), 0) as sum_tvoc,
			COUNT(*) FILTER ( WHERE fire_alarm = 1 ) as fire_alarm_count,
			COUNT(*) FILTER ( WHERE fire_alarm = 0 ) as no_fire_alarm_count
		FROM sensor_readings
		WHERE timestamp >= $1 AND timestamp <= $2
	`

	row := r.db.QueryRow(ctx, query, startTime, endTime)
	return scanSensorStatistics(row)
}

func (r *Repository) Create(ctx context.Context, reading *sensor.Reading) (*uint32, error) {
	const query = `
		INSERT INTO sensor_readings 
		(timestamp, temperature, humidity, tvoc, e_co2, raw_hw, raw_ethanol, pm_25, fire_alarm)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	var id uint32
	err := r.db.QueryRow(ctx, query, sensorReadingArgs(reading)...).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *Repository) Update(ctx context.Context, id uint32, reading *sensor.Reading) error {
	const query = `
		UPDATE sensor_readings
		SET 
		    temperature = $2,
		    humidity = $3,
		    tvoc = $4,
		    e_co2 = $5,
		    raw_hw = $6,
		    raw_ethanol = $7,
		    pm_25 = $8,
		    fire_alarm = $9
		WHERE id = $1
	
	`
	args := sensorReadingArgs(reading)
	args = append([]any{id}, args[1:]...)

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id uint32) error {
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
