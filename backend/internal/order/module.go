package order

import (
	"github.com/0xlakhe/Impluse/internal/app"
	"github.com/0xlakhe/Impluse/internal/payment"
	"github.com/0xlakhe/Impluse/internal/product"
)

type Module struct{
	Handler *Handler
}

func NewModule(application *app.App) *Module{
	orderRepository:=NewRepository(application.DB)

	productRepository:=product.NewRepository(application.DB)

	paymentRepository:=payment.NewRepository(application.DB)

	fakePaymentProvider:=payment.NewFakeProvider()

	paymentService:=payment.NewService(paymentRepository,fakePaymentProvider)

	service:=NewService(orderRepository,productRepository,paymentService)

	handler:=NewHandler(service)

	return &Module{Handler: handler}
}