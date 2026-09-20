package router

import (
	"net/http"

	"github.com/0xlakhe/Impluse/internal/app"
	"github.com/0xlakhe/Impluse/internal/auth"
	"github.com/0xlakhe/Impluse/internal/conversation"
	"github.com/0xlakhe/Impluse/internal/health"
	"github.com/0xlakhe/Impluse/internal/product"
	"github.com/0xlakhe/Impluse/internal/user"
)

func New(application *app.App) http.Handler {
	mux := http.NewServeMux()
	healthModule := health.NewModule(application)
	userModule := user.NewModule(application)
	authModule := auth.NewModule(application)
	productModule := product.NewModule(application)
	conversationModule := conversation.NewModule(application)
	mux.HandleFunc("GET /api/v1/health", healthModule.Handler.Handle)
	mux.HandleFunc("POST /api/v1/auth/register", userModule.Handler.Register)
	mux.HandleFunc("POST /api/v1/auth/login", authModule.Handler.Login)
	mux.Handle("GET /api/v1/auth/me", authModule.Middleware.Authenticate(http.HandlerFunc(authModule.Handler.Me)))
	mux.Handle("GET /api/v1/products", authModule.Middleware.Authenticate(http.HandlerFunc(productModule.Handler.List)))
	mux.Handle("POST /api/v1/sellers/{sellerId}/conversations", authModule.Middleware.Authenticate(http.HandlerFunc(conversationModule.Handler.CreateOrGet)))
	// mux.HandleFunc("GET /api/v1/products",productModule.Handler.List)
	mux.Handle("POST /api/v1/conversations/{conversationId}/messages", authModule.Middleware.Authenticate(http.HandlerFunc(conversationModule.Handler.SendMessage)))
	mux.Handle("GET /api/v1/conversations/{conversationId}/messages", authModule.Middleware.Authenticate(http.HandlerFunc(conversationModule.Handler.GetMessages)))
	return mux
}
