package product

import "github.com/0xlakhe/Impluse/internal/app"

type Module struct {
	Handler *Handler
}

func NewModule(application *app.App) *Module {

	repository := NewRepository(application.DB)
	service := NewService(repository)
	handler := NewHandler(service)

	return &Module{Handler: handler}
}
