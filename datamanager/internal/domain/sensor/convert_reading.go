package sensor

import (
	"github.com/CJovan02/iots/datamanager/protogen/golang/sensorpg"
)

func (r *Reading) ToProto() *sensorpg.GetReadingResponse {
	return &sensorpg.GetReadingResponse{
		Id:          r.Id,
		Timestamp:   r.Timestamp,
		Temperature: r.Temperature,
		Humidity:    r.Humidity,
		Tvoc:        r.Tvoc,
		ECo2:        r.ECo2,
		RawHw:       r.RawHw,
		RawEthanol:  r.RawEthanol,
		Pm_25:       r.PM25,
		FireAlarm:   r.FireAlarm,
	}
}

func ProtoCreateToReading(r *sensorpg.CreateReadingRequest) *Reading {
	return &Reading{
		Timestamp:   r.Timestamp,
		Temperature: r.Temperature,
		Humidity:    r.Humidity,
		Tvoc:        r.Tvoc,
		ECo2:        r.ECo2,
		RawHw:       r.RawHw,
		RawEthanol:  r.RawEthanol,
		PM25:        r.Pm_25,
		FireAlarm:   r.FireAlarm,
	}
}

func ProtoUpdateToReading(r *sensorpg.UpdateReadingRequest) *Reading {
	return &Reading{
		Temperature: r.Temperature,
		Humidity:    r.Humidity,
		Tvoc:        r.Tvoc,
		ECo2:        r.ECo2,
		RawHw:       r.RawHw,
		RawEthanol:  r.RawEthanol,
		PM25:        r.Pm_25,
		FireAlarm:   r.FireAlarm,
	}
}
