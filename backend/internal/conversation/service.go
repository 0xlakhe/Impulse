package conversation

import (
	"context"
	"errors"
	"strings"

	"github.com/0xlakhe/Impluse/internal/ai"
	"github.com/0xlakhe/Impluse/internal/product"
	"github.com/0xlakhe/Impluse/internal/seller"
	"github.com/jackc/pgx/v5"
)

// conv interface
type Repository interface{
	BelongsToUser(ctx context.Context, conversationID string, userID string,)(bool,error)
	
	FindByID(ctx context.Context, conversationID string,)(*Conversation,error)

	CreateMessage(ctx context.Context, conversationID string, role string, content string,)(*Message,error)

	GetMessages(ctx context.Context,conversationID string,)([]Message,error)
	
	CreateOrGet(ctx context.Context, userID string, sellerID string)(*Conversation,error)

	AddProduct(ctx context.Context, conversationID string, productID string,)error
}

//product interface
type ProductRepository interface{
	FindByID(
		ctx context.Context,
		productID string,
	)(*product.Product,error)
}

//seller interface
type SellerRepository interface{
	FindByID(ctx context.Context,sellerID string)(*seller.Seller,error)
}

type Service struct{
	repository Repository
	sellerRepository SellerRepository
	ai ai.Provider
	productRepository ProductRepository
}

func NewService(repository Repository,sellerRepository SellerRepository, aiProvider ai.Provider, productRepository ProductRepository) *Service{
	return &Service{repository: repository,sellerRepository: sellerRepository, ai:aiProvider, productRepository: productRepository}
}

func (s *Service) CreateOrGet(ctx context.Context, userID string, sellerID string, productID string)(*Conversation,error){
	productItem,err:=s.productRepository.FindByID(ctx,productID)
	if err!=nil{
		if errors.Is(err,pgx.ErrNoRows){
			return nil, ErrProductNotFound
		}
		return nil,err
	}

	if productItem.SellerID!=sellerID{
		return nil,ErrProductNotOwnedBySeller
	}
	
	conversation,err:=s.repository.CreateOrGet(ctx,userID,sellerID)

	if err!=nil{
		return nil,err
	}

	if productID!=""{
		err=s.repository.AddProduct(
			ctx,conversation.ID,productID,
		)
		if err!=nil{
			return nil,err
		}
	}
	return conversation,nil
}

func(s *Service) SendMessage(ctx context.Context, userID string,conversationID string, req SendMessageRequest) (*string, error){
	//validate message
	content:=strings.TrimSpace(req.Content)
	if content==""{
		return nil,ErrEmptyMessage
	}

	//make sure conversation belongs to user
	belongs,err:=s.repository.BelongsToUser(ctx,conversationID,userID)
	if err!=nil{
		return nil,err
	}
	if !belongs{
		return nil, ErrConversationNotFound
	}

	// get conversation
	conversation,err:=s.repository.FindByID(
		ctx,conversationID,
	)
	if err!=nil{
		return nil,err
	}

	//get seller
	seller,err:=s.sellerRepository.FindByID(ctx,conversation.SellerID)
	if err!=nil{
		return nil,err
	}

	//get product if one was supplied
	var currentProduct *product.Product

	if req.ProductID!=nil{

		currentProduct,err=s.productRepository.FindByID(ctx,*req.ProductID)
		
		if err!=nil{
			return nil,ErrProductNotFound
		}
		if currentProduct.SellerID!=seller.ID{
			return nil, ErrProductNotOwnedBySeller
		}
	}

	//save user's message
	_,err=s.repository.CreateMessage(
		ctx,conversationID,string(ai.RoleUser),content,
	)
	if err!=nil{
		return nil,err
	}

	//get conversation history
	messages,err:=s.repository.GetMessages(ctx,conversationID)
	if err!=nil{
		return nil,err
	}
	
	//convert database message to ai message
	history:=make([]ai.Message,0,len(messages)+1)
	history=append(history, ai.Message{Role: "system",Content: seller.SystemPrompt})
	for _,message:=range messages{
		history=append(history, ai.Message{Role: message.Role,Content:message.Content},)
	}

	//ask ai for response
	response,err:=s.ai.GenerateResponse(ctx,*seller,currentProduct,history)
	if err!=nil{
		return nil,err
	}

	//save message
	_,err=s.repository.CreateMessage(ctx,conversationID,string(ai.RoleAssistant),*response)
	
	if err!=nil{
		return nil,err
	}
	// return s.repository.GetMessages(ctx,conversationID)
	return response,nil
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
