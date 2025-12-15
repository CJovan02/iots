package grpchand

import (
	"context"

	"github.com/CJovan02/iots/datamanager/internal/domain/sensor"
	"github.com/CJovan02/iots/datamanager/protogen/golang/sensorpg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SensorHandler struct {
	service sensor.Service
	sensorpg.UnimplementedReadingsServer
}

func NewSensorHandler(s sensor.Service) *SensorHandler {
	return &SensorHandler{service: s}
}

func (s *SensorHandler) CountAll(ctx context.Context, _ *emptypb.Empty) (*sensorpg.CountAllResponse, error) {
	count, err := s.service.CountAll(ctx)
	if err != nil {
		return nil, MapErrToGrpc(err)
	}

	response := &sensorpg.CountAllResponse{
		Count: *count,
	}
	return response, nil
}

func (s *SensorHandler) List(ctx context.Context, request *sensorpg.ListReadingsRequest) (*sensorpg.ListReadingsResponse, error) {
	if request.PageSize > 50 {
		return nil, status.Errorf(codes.InvalidArgument, "page size can't be bigger than 50")
	}

	readings, err := s.service.List(ctx, request.PageNumber, request.PageSize)
	if err != nil {
		return nil, MapErrToGrpc(err)
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
	if request.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id must be greater than zero")
	}

	reading, err := s.service.GetById(ctx, request.Id)
	if err != nil {
		return nil, MapErrToGrpc(err)
	}

	return reading.ToProto(), nil
}

func (s *SensorHandler) Statistics(ctx context.Context, request *sensorpg.GetStatisticsRequest) (*sensorpg.GetStatisticsResponse, error) {
	startTime := request.StartTime.AsTime()
	endTime := request.EndTime.AsTime()
	if startTime.IsZero() || endTime.IsZero() || startTime.After(endTime) {
		return nil, status.Errorf(codes.InvalidArgument,
			"start and end time must not be zero and start time must be before end time")
	}

	stat, err := s.service.GetStatistics(ctx, request.StartTime.AsTime(), request.EndTime.AsTime())
	if err != nil {
		return nil, MapErrToGrpc(err)
	}

	return stat.ToProto(), nil
}

func (s *SensorHandler) Create(ctx context.Context, request *sensorpg.CreateReadingRequest) (*sensorpg.CreateReadingResponse, error) {
	reading := sensor.ProtoCreateToReading(request)

	if reading.FireAlarm != 1 && reading.FireAlarm != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "fire alarm must be either 1 or 0")
	}

	if reading.Temperature < -50 || reading.Temperature > 100 {
		return nil, status.Errorf(codes.InvalidArgument, "temperature must be between -50 and 100")
	}
	if reading.Humidity < -50 || reading.Humidity > 100 {
		return nil, status.Errorf(codes.InvalidArgument, "humidity must be between -50 and 100")
	}

	id, err := s.service.Create(ctx, reading)
	if err != nil {
		return nil, MapErrToGrpc(err)
	}

	response := &sensorpg.CreateReadingResponse{Id: *id}
	return response, nil

}
func (s *SensorHandler) Update(ctx context.Context, request *sensorpg.UpdateReadingRequest) (*emptypb.Empty, error) {
	reading := sensor.ProtoUpdateToReading(request)

	if request.Id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id must be greater than zero")
	}

	if reading.FireAlarm != 1 && reading.FireAlarm != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "fire alarm must be either 1 or 0")
	}

	if reading.Temperature < -50 || reading.Temperature > 100 {
		return nil, status.Errorf(codes.InvalidArgument, "temperature must be between -50 and 100")
	}
	if reading.Humidity < -50 || reading.Humidity > 100 {
		return nil, status.Errorf(codes.InvalidArgument, "humidity must be between -50 and 100")
	}

	err := s.service.Update(ctx, request.Id, reading)
	if err != nil {
		return nil, MapErrToGrpc(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *SensorHandler) Delete(ctx context.Context, request *sensorpg.DeleteReadingRequest) (*emptypb.Empty, error) {
	if request.Id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id must be greater than zero")
	}

	err := s.service.Delete(ctx, request.Id)
	if err != nil {
		return nil, MapErrToGrpc(err)
	}

	return &emptypb.Empty{}, nil
}
