package grpchand

import (
	"context"
	"fmt"

	"github.com/CJovan02/iots/datamanager/internal/domain/sensor"
	"github.com/CJovan02/iots/datamanager/protogen/golang/sensorpg"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SensorHandler struct {
	service sensor.Service
	sensorpg.UnimplementedReadingsServer
}

func NewSensorHandler(s sensor.Service) *SensorHandler {
	return &SensorHandler{service: s}
}

func (s *SensorHandler) CountAll(ctx context.Context,
	_ *emptypb.Empty,
) (*sensorpg.CountAllResponse, error) {
	count, err := s.service.CountAll(ctx)
	if err != nil {
		return nil, err
	}

	response := &sensorpg.CountAllResponse{
		Count: *count,
	}
	return response, nil
}

func (s *SensorHandler) List(ctx context.Context,
	request *sensorpg.ListReadingsRequest,
) (*sensorpg.ListReadingsResponse, error) {
	if request.PageSize > 50 {
		return nil, sensor.NewInvalidArgument("page size can't be bigger than 50")
	}

	readings, err := s.service.List(ctx, request.PageNumber, request.PageSize)
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

func (s *SensorHandler) Get(
	ctx context.Context,
	request *sensorpg.GetReadingRequest,
) (*sensorpg.GetReadingResponse, error) {
	if request.Id == 0 {
		return nil, sensor.NewInvalidArgument("id must be greater than zero")
	}

	reading, err := s.service.GetById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return reading.ToProto(), nil
}

func (s *SensorHandler) Statistics(ctx context.Context,
	request *sensorpg.GetStatisticsRequest,
) (*sensorpg.GetStatisticsResponse, error) {
	startTime := request.StartTime.AsTime()
	endTime := request.EndTime.AsTime()
	if startTime.IsZero() || endTime.IsZero() || startTime.After(endTime) {
		return nil, sensor.NewInvalidArgument(
			"start and end time must not be zero and start time must be before end time",
		)
	}

	stat, err := s.service.GetStatistics(ctx, request.StartTime.AsTime(), request.EndTime.AsTime())
	if err != nil {
		return nil, err
	}

	return stat.ToProto(), nil
}

func (s *SensorHandler) Create(
	ctx context.Context,
	request *sensorpg.CreateReadingRequest,
) (*sensorpg.CreateReadingResponse, error) {
	reading := sensor.ProtoCreateToReading(request)

	err := ValidateReading(reading)
	if err != nil {
		return nil, err
	}

	id, err := s.service.Create(ctx, reading)
	if err != nil {
		return nil, err
	}

	response := &sensorpg.CreateReadingResponse{Id: *id}
	return response, nil

}

func (s *SensorHandler) BatchCreate(
	ctx context.Context,
	request *sensorpg.BatchCreateReadingsRequest,
) (*sensorpg.BatchCreateReadingsResponse, error) {
	if len(request.ReadingRequests) == 0 || len(request.ReadingRequests) > 100 {
		return nil, sensor.NewInvalidArgument(
			"invalid number of readings, it must be between 1 and 100",
		)
	}

	readings := make([]*sensor.Reading, 0, len(request.ReadingRequests))

	for index, reqReading := range request.ReadingRequests {
		reading := sensor.ProtoCreateToReading(reqReading)
		readings = append(readings, reading)

		err := ValidateReadingWithIndex(reading, &index)
		if err != nil {
			return nil, err
		}
	}

	ids, err := s.service.BatchCreate(ctx, readings)
	if err != nil {
		return nil, err
	}

	return &sensorpg.BatchCreateReadingsResponse{Ids: ids}, nil
}

func (s *SensorHandler) Update(
	ctx context.Context,
	request *sensorpg.UpdateReadingRequest,
) (*emptypb.Empty, error) {
	reading := sensor.ProtoUpdateToReading(request)

	if request.Id <= 0 {
		return nil, sensor.NewInvalidArgument("id must be greater than zero")
	}

	err := ValidateReading(reading)
	if err != nil {
		return nil, err
	}

	err = s.service.Update(ctx, request.Id, reading)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SensorHandler) Delete(ctx context.Context,
	request *sensorpg.DeleteReadingRequest,
) (*emptypb.Empty, error) {
	if request.Id <= 0 {
		return nil, sensor.NewInvalidArgument("id must be greater than zero")
	}

	err := s.service.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func ValidateReading(reading *sensor.Reading) error {
	return validateReading(reading, false, nil)
}

func ValidateReadingWithIndex(reading *sensor.Reading, index *int) error {
	return validateReading(reading, true, index)
}

func validateReading(reading *sensor.Reading, printIndex bool, index *int) error {
	if reading.FireAlarm != 1 && reading.FireAlarm != 0 {
		var msg string
		if printIndex {
			msg = fmt.Sprintf("%d: ", *index)
		}
		msg += "fire alarm must be either 1 or 0"
		return sensor.NewInvalidArgument(msg)
	}

	if reading.Temperature < -50 || reading.Temperature > 100 {
		var msg string
		if printIndex {
			msg = fmt.Sprintf("%d: ", *index)
		}
		msg += "temperature must be between -50 and 100"
		return sensor.NewInvalidArgument(msg)
	}
	if reading.Humidity < -50 || reading.Humidity > 100 {
		var msg string
		if printIndex {
			msg = fmt.Sprintf("%d: ", *index)
		}
		msg += "humidity must be between -50 and 100"
		return sensor.NewInvalidArgument(msg)
	}
	return nil
}
