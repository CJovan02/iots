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
