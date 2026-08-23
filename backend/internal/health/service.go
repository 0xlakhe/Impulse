package health

import (
	"context"

)

type Service struct{
	repository *Repository
}

func NewService(repository *Repository,) *Service{
	return &Service{
		repository: repository,
	}
}

func (s *Service) CheckDatabase(ctx context.Context) error{
	return s.repository.Ping(ctx)
}