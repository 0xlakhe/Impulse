package ai

import (
	"context"
	"fmt"

	"github.com/0xlakhe/Impluse/internal/seller"
)

type FakeProvider struct{
}

func NewFakeProvider() *FakeProvider{
	return &FakeProvider{}
}

func (f *FakeProvider) GenerateResponse(ctx context.Context,seller seller.Seller, history []Message,)(string,error){
	if len(history)==0{
		return fmt.Sprintf("Hey! I'm %s, What would you like to know?",seller.Name),nil
	}
	lastMessage:=history[len(history)-1]
	return fmt.Sprintf("%s here! You asked:\"%s\". That's greate dudubutter",seller.Name,lastMessage.Content),nil
}