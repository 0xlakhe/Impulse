package auth

import (
	"context"

	"github.com/0xlakhe/Impluse/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{
	users *user.Repository
	jwt *JWTManager
}

func NewService(users *user.Repository,jwt *JWTManager) *Service{
	return &Service{users: users,jwt: jwt}
}

func (s *Service) Login(ctx context.Context, req LoginRequest)(*LoginResponse,error){
	u,err:=s.users.FindByEmail(ctx,req.Email)
	if err!=nil{
		return nil, user.ErrInvalidCredentials
	}

	err=bcrypt.CompareHashAndPassword(
		[]byte(u.PasswordHash),
		[]byte(req.Password),
	)
	if err!=nil{
		return nil, user.ErrInvalidCredentials
	}

	token,err:=s.jwt.Generate(u.ID)
	if err!=nil{
		return nil, err
	}
	return &LoginResponse{Token: token},nil
}