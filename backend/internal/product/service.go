package product

import "context"

type Service struct{
	repository *Repository
}

func NewService (repository *Repository) *Service{
	return &Service{repository: repository}
}

func(s *Service)List(ctx context.Context,)([]ProductResponse,error){
	return s.repository.List(ctx,)
}