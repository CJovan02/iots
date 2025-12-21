package sensorsvc

import (
	"context"

	"github.com/CJovan02/iots/datamanager/internal/domain/sensor"
)

type Service struct {
	repo      sensor.Repository
	publisher sensor.ReadingsPublisher
	topic     string
}

func New(repo sensor.Repository, publisher sensor.ReadingsPublisher, topic string) *Service {
	return &Service{repo: repo, publisher: publisher, topic: topic}
}

func (s *Service) CountAll(ctx context.Context) (*uint32, error) {
	return s.repo.CountAll(ctx)
}

func (s *Service) List(ctx context.Context, pageNumber uint32, pageSize uint32) ([]sensor.Reading, error) {
	offset := (pageNumber - 1) * pageSize

	return s.repo.List(ctx, offset, pageSize)
}

func (s *Service) GetById(ctx context.Context, id uint32) (*sensor.Reading, error) {
	return s.repo.GetById(ctx, id)
}

func (s *Service) GetStatistics(ctx context.Context, startTime int64, endTime int64) (*sensor.Statistics, error) {
	return s.repo.GetStatistics(ctx, startTime, endTime)
}

func (s *Service) Create(ctx context.Context, reading *sensor.Reading) (*uint32, error) {
	id, err := s.repo.Create(ctx, reading)
	if err != nil {
		return nil, err
	}

	// publish reading to MQTT topic
	err = s.publisher.PublishJson(s.topic, reading)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *Service) BatchCreate(ctx context.Context, readings []*sensor.Reading) ([]uint32, error) {
	ids, err := s.repo.BatchCreate(ctx, readings)
	if err != nil {
		return nil, err
	}

	// publish reading to MQTT broker
	err = s.publisher.PublishJson(s.topic, readings)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (s *Service) Update(ctx context.Context, id uint32, reading *sensor.Reading) error {
	exists, err := s.repo.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		return sensor.NewNotFound(id)
	}

	return s.repo.Update(ctx, id, reading)
}

func (s *Service) Delete(ctx context.Context, id uint32) error {
	return s.repo.Delete(ctx, id)
}
