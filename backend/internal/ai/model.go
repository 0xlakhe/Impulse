package ai

import (
	"context"

	"github.com/0xlakhe/Impluse/internal/product"
	"github.com/0xlakhe/Impluse/internal/seller"
)

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Category    string
}

type Provider interface {
	GenerateResponse(
		ctx context.Context,
		seller seller.Seller,
		product *product.Product,
		history []Message,
	) (*string, error)
}
