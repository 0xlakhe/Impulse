package auth

import (
	"github.com/0xlakhe/Impluse/internal/app"
	"github.com/0xlakhe/Impluse/internal/user"
)


type Module struct{
	Handler *Handler
	Middleware *Middleware
}


func NewModule(application *app.App) *Module{
	userRepo:=user.NewRepository(application.DB)
	jwtManager:=NewJWTManager(application.Config.JWTSecret)
	service:=NewService(userRepo,jwtManager)
	middleware:=NewMiddleware(jwtManager)
	handler:=NewHandler(service)
	
	return &Module{Handler: handler,Middleware: middleware}
}