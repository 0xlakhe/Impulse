package ai

import (
	"context"

	"github.com/0xlakhe/Impluse/internal/seller"
)

const (
	RoleUser="user"
	RoleAssistant="assistant"
)

type Message struct{
	Role string
	Content string
}

type Provider interface{
	GenerateResponse(
		ctx context.Context,
		seller seller.Seller,
		history []Message,
	)(string,error)
}

