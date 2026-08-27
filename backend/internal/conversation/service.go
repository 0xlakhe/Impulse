package conversation

import (
	"context"
	"strings"
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

func(s *Service) SendMessage(ctx context.Context, conversationID string, userID string, content string) (*Message, error){
	
	exists,err:=s.repository.BelongsToUser(
		ctx,conversationID,userID,
	)
	content=strings.TrimSpace(content)
	if content==""{
		return nil,ErrEmptyMessage
	}

	if err!=nil{
		return nil,err
	}
	if !exists{
		return nil, ErrConversationNotFound
	}

	return s.repository.CreateMessage(ctx,conversationID,"user",content)
}

func (s *Service) GetMessags(ctx context.Context,conversationID string, userID string,)([]Message, error){
	exists,err:=s.repository.BelongsToUser(
		ctx,conversationID,userID,
	)
	if err!=nil{
		return nil,err
	}
	if !exists{
		return nil, ErrConversationNotFound
	}
	return s.repository.GetMessages(ctx,conversationID,)
}