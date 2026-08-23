package router

import (
	"net/http"

	"github.com/0xlakhe/Impluse/internal/app"
	"github.com/0xlakhe/Impluse/internal/auth"
	"github.com/0xlakhe/Impluse/internal/health"
	"github.com/0xlakhe/Impluse/internal/user"
)

func New(application *app.App) http.Handler{
	mux:=http.NewServeMux()
	healthModule:=health.NewModule(application)
	userModule:=user.NewModule(application)
	authModule:=auth.NewModule(application)
	mux.HandleFunc("GET /api/v1/health",healthModule.Handler.Handle)
	mux.HandleFunc("POST /api/v1/auth/register",userModule.Handler.Register)
	mux.HandleFunc("POST /api/v1/auth/login",authModule.Handler.Login)
	mux.Handle("GET /api/v1/auth/me",authModule.Middleware.Authenticate(http.HandlerFunc(authModule.Handler.Me)))	
	return mux
}