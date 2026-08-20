package health

import (
	"github.com/0xlakhe/Impluse/internal/app"
	healthrepo "github.com/0xlakhe/Impluse/internal/health/repository"
)


type Module struct{
	Handler *Handler
}

func NewModule(application *app.App) *Module{
	repository:=healthrepo.New(application.DB)
	service:=NewService(repository)
	handler:=NewHandler(service)
	return &Module{Handler: handler}
}