package sensorsvc

import (
	"context"
	"time"

	"github.com/CJovan02/iots/datamanager/internal/domain/sensor"
	"github.com/CJovan02/iots/datamanager/internal/reading_errors"
)

type Service struct {
	repo sensor.Repository
}

func New(repo sensor.Repository) *Service {
	return &Service{repo: repo}
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

func (s *Service) GetStatistics(ctx context.Context, startTime time.Time, endTime time.Time) (*sensor.Statistics, error) {
	return s.repo.GetStatistics(ctx, startTime, endTime)
}

func (s *Service) Create(ctx context.Context, reading *sensor.Reading) (*uint32, error) {
	return s.repo.Create(ctx, reading)
}

func (s *Service) BatchCreate(ctx context.Context, readings []*sensor.Reading) ([]uint32, error) {
	return s.repo.BatchCreate(ctx, readings)
}

func (s *Service) Update(ctx context.Context, id uint32, reading *sensor.Reading) error {
	exists, err := s.repo.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		return reading_errors.NewNotFound("reading", id)
	}

	return s.repo.Update(ctx, id, reading)
}

func (s *Service) Delete(ctx context.Context, id uint32) error {
	return s.repo.Delete(ctx, id)
}
