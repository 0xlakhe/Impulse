package conversation

import (
	"github.com/0xlakhe/Impluse/internal/ai"
	"github.com/0xlakhe/Impluse/internal/app"
	"github.com/0xlakhe/Impluse/internal/product"
	"github.com/0xlakhe/Impluse/internal/seller"
)

type Module struct{
	Handler *Handler
}

func NewModule(application *app.App) *Module{
	repository:=NewRepository(application.DB)
	sellerRepository:=seller.NewRepository(application.DB)
	aiProvider:=ai.NewFakeProvider()
	productRepository:=product.NewRepository(application.DB)
	service:=NewService(repository,sellerRepository,aiProvider,productRepository)
	handler:=NewHandler(service)
	return &Module{Handler: handler}
}