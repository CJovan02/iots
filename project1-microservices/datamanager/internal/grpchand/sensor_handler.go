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

func (s *SensorHandler) Get(ctx context.Context, request *sensorpg.GetReadingRequest) (*sensorpg.GetReadingResponse, error) {
	reading, err := s.service.GetById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return reading.ToProto(), nil
}

func (s *SensorHandler) Statistics(ctx context.Context, request *sensorpg.GetStatisticsRequest) (*sensorpg.GetStatisticsResponse, error) {
	stat, err := s.service.GetStatistics(ctx, request.StartTime.AsTime(), request.EndTime.AsTime())
	if err != nil {
		return nil, err
	}

	return stat.ToProto(), nil
}

func (s *SensorHandler) Create(ctx context.Context, request *sensorpg.CreateReadingRequest) (*sensorpg.CreateReadingResponse, error) {
	reading := sensor.ProtoCreateToReading(request)

	id, err := s.service.Create(ctx, reading)
	if err != nil {
		return nil, err
	}

	response := &sensorpg.CreateReadingResponse{Id: *id}
	return response, nil

}
func (s *SensorHandler) Update(ctx context.Context, request *sensorpg.UpdateReadingRequest) (*emptypb.Empty, error) {
	reading := sensor.ProtoUpdateToReading(request)

	err := s.service.Update(ctx, request.Id, reading)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SensorHandler) Delete(ctx context.Context, request *sensorpg.DeleteReadingRequest) (*emptypb.Empty, error) {
	err := s.service.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
