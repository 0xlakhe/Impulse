package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service struct{
	repository *Repository
}

func NewService(
	repository *Repository,
)*Service{
	return &Service{repository: repository}
}

func(s *Service) Register(ctx context.Context, req RegisterRequest)(*RegisterResponse, error){
	if err:=ValidateRegisterRequest(req); err!=nil{
		return nil, err
	}

	//checking if email exists
	existingUser,err:=s.repository.FindByEmail(ctx,req.Email)
	if err==nil && existingUser!=nil{
		return nil, ErrEmailAlreadyExists
	}
	if err!=nil && !IsNotFound(err){
		return nil,err
	}
	//checking is username exists
	existingUser,err=s.repository.FindByUsername(ctx,req.Username)
	if err==nil && existingUser!=nil{
		return nil, ErrUserNameTaken
	}
	if err!=nil && !IsNotFound(err){
		return nil,err
	}

	//password hashing
	passwordHash,err:=bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err!=nil{
		return nil,err
	}
	user:=&User{
		Username: req.Username,
		Email: req.Email,
		PasswordHash: string(passwordHash),
	}
	if err:=s.repository.Create(ctx,user); err!=nil{
		return nil,err
	}
	return &RegisterResponse{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
	},nil
}