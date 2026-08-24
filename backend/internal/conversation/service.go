package conversation

import (
	"context"

)

type Service struct{
	repository *Repository
}

func NewService(repository *Repository) *Service{
	return &Service{repository: repository}
}

func (s *Service) CreateOrGet(ctx context.Context, userID string, sellerID string)(*Conversation,error){

	return s.repository.CreateOrGet(ctx,userID,sellerID)
}