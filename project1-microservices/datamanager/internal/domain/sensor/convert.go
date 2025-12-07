package sensor

import (
	"github.com/CJovan02/iots/project1-microservices/datamanager/protogen/golang/sensorpg"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *Reading) ToProto() *sensorpg.GetReadingResponse {
	return &sensorpg.GetReadingResponse{
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
