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

func (f *FakeProvider) GenerateResponse(ctx context.Context,seller seller.Seller,p *product.Product, history []Message,)(*string,error){
	prompt:=BuildPrompt(seller,p,history)
	if len(history)==0{
		res:=fmt.Sprintf("Fake AI received this prompt:\n\n%s",prompt)
		return &res,nil
	}
	lastMessage:=history[len(history)-1]

	res:=fmt.Sprintf("%s here! You asked:\"%s\". That's greate dudubutter",seller.Name,lastMessage.Content)
	return &res,nil
}