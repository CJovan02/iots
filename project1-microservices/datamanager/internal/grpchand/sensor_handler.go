package grpchand

import (
	"context"

	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/domain/sensor"
	"github.com/CJovan02/iots/project1-microservices/datamanager/protogen/golang/sensorpg"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SensorHandler struct {
	service sensor.Service
	sensorpg.UnimplementedReadingsServer
}

func NewSensorHandler(s sensor.Service) *SensorHandler {
	return &SensorHandler{service: s}
}

func (s *SensorHandler) List(ctx context.Context, _ *emptypb.Empty) (*sensorpg.ListReadingsResponse, error) {
	readings, err := s.service.List(ctx)
	if err != nil {
		return nil, err
	}

	var readingsP []*sensorpg.GetReadingResponse
	for _, reading := range readings {
		readingsP = append(readingsP, reading.ToProto())
	}

	response := &sensorpg.ListReadingsResponse{
		Readings: readingsP,
	}
	return response, nil
}

//Get(context.Context, *GetReadingRequest) (*GetReadingResponse, error)
//Create(context.Context, *CreateReadingRequest) (*CreateReadingResponse, error)
//Update(context.Context, *UpdateReadingRequest) (*UpdateReadingResponse, error)
//Delete(context.Context, *DeleteReadingRequest) (*DeleteReadingResponse, error)
