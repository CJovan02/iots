package sensor

import (
	"github.com/CJovan02/iots/datamanager/protogen/golang/sensorpg"
)

func (r *Reading) ToProto() *sensorpg.GetReadingResponse {
	return &sensorpg.GetReadingResponse{
		Id:          r.Id,
		DeviceId:    r.DeviceId,
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
	return NewReading(
		0,
		*r.DeviceId,
		r.Timestamp,
		r.Temperature,
		r.Humidity,
		r.Tvoc,
		r.ECo2,
		r.RawHw,
		r.RawEthanol,
		r.Pm_25,
		r.FireAlarm)
}

func ProtoUpdateToReading(r *sensorpg.UpdateReadingRequest) *Reading {
	return NewReading(
		0,
		"",
		0,
		r.Temperature,
		r.Humidity,
		r.Tvoc,
		r.ECo2,
		r.RawHw,
		r.RawEthanol,
		r.Pm_25,
		r.FireAlarm)

}
