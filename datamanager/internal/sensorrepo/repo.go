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
			MIN(temperature) as min_temp,
			MAX(temperature) as max_temp,
			AVG(temperature) as avg_temp,
			MIN(humidity) as min_humidity,
			MAX(humidity) as max_humidity,
			AVG(humidity) as avg_humidity,
			SUM(tvoc) as sum_tvoc,
			COUNT(*) FILTER ( WHERE fire_alarm = 1 ) as fire_alarm_count,
			COUNT(*) FILTER ( WHERE fire_alarm = 0 ) as no_fire_alarm_count
		FROM sensor_readings
		where timestamp >= $1 and timestamp <= $2
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
