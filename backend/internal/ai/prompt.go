package ai

import (
	"fmt"

	"github.com/0xlakhe/Impluse/internal/product"
	"github.com/0xlakhe/Impluse/internal/seller"
)


func BuildPrompt(s seller.Seller, p *product.Product, history []Message) string{
	prompt:=fmt.Sprintf(`
	You are %s.
	PERSONA:
	%s
	INSTRUCTIONS:
	%s
	`,s.Name,s.Persona,s.SystemPrompt)

	if p!=nil{
		prompt+=fmt.Sprintf(`
		CURRENT PRODUCT:
		Name: %s
		Category: %s
		Price: %.2f
		Description: %s
		`,p.Name,p.Category,p.Price,p.Description)
	}
	prompt+="\nCONVERSATION HISTORY:\n"

	for _,message:=range history{
		prompt+=fmt.Sprintf(
			"%s: %s\n",message.Role,message.Content)
	}
	return prompt
}

func BuildRequestMessages(s seller.Seller,p *product.Product,history []Message)[]Message{
	var messages []Message

	systemContent:=fmt.Sprintf(
		"You are %s.\n PERSONA:\n%sINSTRUCTIONS:\n%s",s.Name,s.Persona,s.SystemPrompt,
	)
	
	if p!=nil{
		systemContent=fmt.Sprintf("\n\nCURRENT PRODUCT:\nName:%s\nCategory:%s\nPrice: %.2f\nDescription: %s",p.Name,p.Category,p.Price,p.Description)
		fmt.Println(p.Name,p.Category,p.Price,p.Description)
	}
	messages=append(messages, Message{
		Role: string(RoleSystem),
		Content: systemContent,
	})

	messages=append(messages, history...)
	return messages
}