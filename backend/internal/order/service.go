package order

import (
	"context"
	"fmt"
	"time"

	"github.com/0xlakhe/Impluse/internal/product"
	"github.com/google/uuid"
)

type ProductRepository interface {
	FindByID(ctx context.Context, productID string) (*product.Product, error)
}

type Repository interface {
	Create(ctx context.Context, order *Order, items []OrderItem) error
}

type Service struct {
	repository        Repository
	productRepository ProductRepository
}

func NewService(repository Repository, productRepository ProductRepository) *Service {
	return &Service{repository: repository, productRepository: productRepository}
}

func (s *Service) Create(ctx context.Context, userID string, req CreateOrderRequest) (*Order, error) {
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
	return order, nil
}
