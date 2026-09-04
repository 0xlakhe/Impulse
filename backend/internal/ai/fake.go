package ai

import (
	"context"
	"fmt"

	"github.com/0xlakhe/Impluse/internal/product"
	"github.com/0xlakhe/Impluse/internal/seller"
)

type FakeProvider struct{
}

func NewFakeProvider() *FakeProvider{
	return &FakeProvider{}
}

func (f *FakeProvider) GenerateResponse(ctx context.Context,seller seller.Seller,p *product	.Product, history []Message,)(string,error){
	if len(history)==0{
		return fmt.Sprintf("Hey! I'm %s, What would you like to know?",seller.Name),nil
	}
	lastMessage:=history[len(history)-1]

	if p!=nil{
		return fmt.Sprintf("%s: You're asking about %s. You said: \"%s\"",seller.Name,p.Name,lastMessage.Content),nil
	}
	return fmt.Sprintf("%s here! You asked:\"%s\". That's greate dudubutter",seller.Name,lastMessage.Content),nil
}