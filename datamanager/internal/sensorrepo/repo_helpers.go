package sensorrepo

import (
	"github.com/CJovan02/iots/datamanager/internal/domain/sensor"
	"github.com/jackc/pgx/v5"
)

func scanSensorReading(row pgx.Row) (*sensor.Reading, error) {
	var sr sensor.Reading

	err := row.Scan(
		&sr.Id,
		&sr.Timestamp,
		&sr.Temperature,
		&sr.Humidity,
		&sr.Tvoc,
		&sr.ECo2,
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

func scanSensorReadingValue(row pgx.CollectableRow) (sensor.Reading, error) {
	reading, err := scanSensorReading(row)
	if err != nil {
		return sensor.Reading{}, err
	}
	return *reading, err
}

func sensorReadingArgs(r *sensor.Reading) []any {
	return []any{
		r.Timestamp,
		r.Temperature,
		r.Humidity,
		r.Tvoc,
		r.ECo2,
		r.RawHw,
		r.RawEthanol,
		r.PM25,
		r.FireAlarm,
	}
}

func scanSensorStatistics(row pgx.Row) (*sensor.Statistics, error) {
	var s sensor.Statistics

	err := row.Scan(
		&s.ReadingsCount,
		&s.MinTemperature,
		&s.MaxTemperature,
		&s.AvgTemperature,
		&s.MinHumidity,
		&s.MaxHumidity,
		&s.AvgHumidity,
		&s.SumTVOC,
		&s.FireAlarmCount,
		&s.NoFireAlarmCount,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
