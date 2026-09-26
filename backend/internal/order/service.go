package order

import (
	"context"
	"fmt"
	"time"

	"github.com/0xlakhe/Impluse/internal/payment"
	"github.com/0xlakhe/Impluse/internal/product"
	"github.com/google/uuid"
)

type ProductRepository interface {
	FindByID(ctx context.Context, productID string) (*product.Product, error)
}

type PaymentProcessor interface{
	Charge(ctx context.Context, orderID string, amount float64, idempotencyKey string)(*payment.Payment,error)
}

type Service struct {
	repository        Repository
	productRepository ProductRepository
	paymentProcessor PaymentProcessor
}

func NewService(repository Repository, productRepository ProductRepository, paymentProcessor PaymentProcessor) *Service {
	return &Service{repository: repository, productRepository: productRepository, paymentProcessor: paymentProcessor}
}

func (s *Service) Create(ctx context.Context, userID string, req CreateOrderRequest, idempotencyKey string) (*Order, error) {
	if len(req.Items) == 0 {
		return nil, ErrEmptyOrders
	}

	orderItems := make([]OrderItem, 0, len(req.Items))

	var total float64

	now := time.Now()

	order := &Order{
		ID:        uuid.NewString(),
		UserID:    userID,
		Status:    StatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}

	for _, item := range req.Items {
		if item.Quantity <= 0 {
			return nil, ErrInvalidQuantity
		}
		productItem, err := s.productRepository.FindByID(ctx, item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrProductNotFound, item.ProductID)
		}

		subtotal := productItem.Price * float64(item.Quantity)

		orderItems = append(orderItems, OrderItem{
			ID:          uuid.NewString(),
			OrderID:     order.ID,
			ProductID:   productItem.ID,
			ProductName: productItem.Name,
			UnitPrice:   productItem.Price,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})

		total += subtotal
	}
	order.TotalAmount = total

	err := s.repository.Create(ctx, order, orderItems)

	if err != nil {
		return nil, err
	}

	paymentResult,err:=s.paymentProcessor.Charge(ctx,order.ID,order.TotalAmount,idempotencyKey)

	if err!=nil{
		_=s.repository.UpdateStatus(ctx, order.ID,StatusFailed)

		order.Status=StatusFailed

		return order,err
	}

	switch paymentResult.Status{
	case payment.StatusSuccess:
		err=s.repository.UpdateStatus(ctx,order.ID,StatusPaid)
		if err!=nil{
			return order,err
		}
		order.Status=StatusPaid
	case payment.StatusFailed:
		_=s.repository.UpdateStatus(ctx,order.ID,StatusFailed)
		order.Status=StatusFailed
		return order,payment.ErrPaymentFailed
	default:
		order.Status=StatusPending
	}

	return order, nil
}
