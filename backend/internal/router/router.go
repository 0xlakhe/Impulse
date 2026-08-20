package router

import (
	"net/http"

	"github.com/0xlakhe/Impluse/internal/app"
	"github.com/0xlakhe/Impluse/internal/health"
)

func New(application *app.App) http.Handler{
	mux:=http.NewServeMux()
	healthModule:=health.NewModule(application)
	mux.HandleFunc("/api/v1/health",healthModule.Handler.Handle)
	return mux
}