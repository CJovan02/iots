package sensorsvc

import (
	"context"
	"errors"

	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/domain/sensor"
)

type Service struct {
	repo sensor.Repository
}

func New(repo sensor.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetById(ctx context.Context, id int32) (*sensor.Reading, error) {
	if id < 0 {
		return nil, errors.New("id must be positive")
	}

	return s.repo.GetById(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]sensor.Reading, error) {
	return s.repo.List(ctx)
}

func (s *Service) Create(ctx context.Context, reading *sensor.Reading) error {
	if reading.FireAlarm != 1 && reading.FireAlarm != 0 {
		return errors.New("fire alarm must be either 1 or 0")
	}

	if reading.Temperature < -50 || reading.Temperature > 100 {
		return errors.New("temperature must be between 0 and 100")
	}
	if reading.Humidity < -50 || reading.Humidity > 100 {
		return errors.New("humidity must be between 0 and 100")
	}

	return s.repo.Create(ctx, reading)
}

func (s *Service) Update(ctx context.Context, id int32, reading *sensor.Reading) error {
	if id < 0 {
		return errors.New("id must be positive")
	}

	if reading.FireAlarm != 1 && reading.FireAlarm != 0 {
		return errors.New("fire alarm must be either 1 or 0")
	}

	if reading.Temperature < -50 || reading.Temperature > 100 {
		return errors.New("temperature must be between 0 and 100")
	}
	if reading.Humidity < -50 || reading.Humidity > 100 {
		return errors.New("humidity must be between 0 and 100")
	}
	return s.repo.Update(ctx, id, reading)
}

func (s *Service) Delete(ctx context.Context, id int32) error {
	if id < 0 {
		return errors.New("id must be positive")
	}

	return s.repo.Delete(ctx, id)
}
