package server

import (
	"net/http"

	"github.com/0xlakhe/Impluse/internal/app"
	"github.com/0xlakhe/Impluse/internal/database"
	"github.com/0xlakhe/Impluse/internal/middleware"
	"github.com/0xlakhe/Impluse/internal/router"
)

type Server struct {
	httpServer *http.Server
	db         *database.Database
}

func New(application *app.App) *Server {
	handler := router.New(application)
	handler = middleware.Logging(handler)

	httpServer := &http.Server{
		Addr:    ":" + application.Config.Port,
		Handler: handler,
	}
	return &Server{
		httpServer: httpServer,
		db:         application.DB,
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
