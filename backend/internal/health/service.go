package health

import (
	"context"

	"github.com/0xlakhe/Impluse/internal/health/repository"
)

type Service struct{
	repository *repository.Repository
}

func NewService(repository *repository.Repository,) *Service{
	return &Service{
		repository: repository,
	}
}

func (s *Service) CheckDatabase(ctx context.Context) error{
	return s.repository.Ping(ctx)
}