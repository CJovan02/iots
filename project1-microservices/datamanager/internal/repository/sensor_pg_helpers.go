package repository

import (
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/domain/sensor"
	"github.com/jackc/pgx/v5"
)

func scanSensorReading(row pgx.Row) (*sensor.SensorReading, error) {
	var sr sensor.SensorReading

	err := row.Scan(
		&sr.Id,
		&sr.Timestamp,
		&sr.Temperature,
		&sr.Humidity,
		&sr.TVOC,
		&sr.ECO2,
		&sr.RawHw,
		&sr.RawEthanol,
		&sr.PM25,
		&sr.FireAlarm,
	)
	if err != nil {
		return nil, err
	}
	return &sr, nil
}

func scanSensorReadingValue(row pgx.CollectableRow) (sensor.SensorReading, error) {
	reading, err := scanSensorReading(row)
	if err != nil {
		return sensor.SensorReading{}, err
	}
	return *reading, err
}

func sensorReadingArgs(r *sensor.SensorReading) []any {
	return []any{
		r.Timestamp,
		r.Temperature,
		r.Humidity,
		r.TVOC,
		r.ECO2,
		r.RawHw,
		r.RawEthanol,
		r.PM25,
		r.FireAlarm,
	}
}
