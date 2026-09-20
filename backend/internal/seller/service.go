package seller

import "context"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) FindByID(ctx context.Context, sellerID string) (*Seller, error) {
	return s.repository.FindByID(ctx, sellerID)
}
