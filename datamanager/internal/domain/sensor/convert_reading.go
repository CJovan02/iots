package sensor

import (
	"github.com/CJovan02/iots/datamanager/protogen/golang/sensorpg"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *Reading) ToProto() *sensorpg.GetReadingResponse {
	return &sensorpg.GetReadingResponse{
		Id:          r.Id,
		Timestamp:   timestamppb.New(r.Timestamp),
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

func ProtoCreateToReading(request *sensorpg.CreateReadingRequest) *Reading {
	return &Reading{
		Timestamp:   request.Timestamp.AsTime(),
		Temperature: request.Temperature,
		Humidity:    request.Humidity,
		Tvoc:        request.Tvoc,
		ECo2:        request.ECo2,
		RawHw:       request.RawHw,
		RawEthanol:  request.RawEthanol,
		PM25:        request.Pm_25,
		FireAlarm:   request.FireAlarm,
	}
}

func ProtoUpdateToReading(request *sensorpg.UpdateReadingRequest) *Reading {
	return &Reading{
		Temperature: request.Temperature,
		Humidity:    request.Humidity,
		Tvoc:        request.Tvoc,
		ECo2:        request.ECo2,
		RawHw:       request.RawHw,
		RawEthanol:  request.RawEthanol,
		PM25:        request.Pm_25,
		FireAlarm:   request.FireAlarm,
	}
}
