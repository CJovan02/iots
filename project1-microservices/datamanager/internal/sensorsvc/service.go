package sensorsvc

import (
	"context"
	"time"

	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/domain/sensor"
)

type Service struct {
	repo sensor.Repository
}

func New(repo sensor.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context) ([]sensor.Reading, error) {
	return s.repo.List(ctx)
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

func (s *Service) Update(ctx context.Context, id uint32, reading *sensor.Reading) error {
	return s.repo.Update(ctx, id, reading)
}

func (s *Service) Delete(ctx context.Context, id uint32) error {
	return s.repo.Delete(ctx, id)
}
