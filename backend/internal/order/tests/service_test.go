package tests

import (
	"context"
	"testing"

	"github.com/0xlakhe/Impluse/internal/order"
	"github.com/0xlakhe/Impluse/internal/product"
)

type fakeRepository struct {
	createdOrder *order.Order
	createdItems []order.OrderItem
}

func (f *fakeRepository) Create(ctx context.Context, order *order.Order, items []order.OrderItem) error {
	f.createdOrder = order
	f.createdItems = items
	return nil
}

type fakeProductRepository struct {
	products map[string]*product.Product
}

func (f *fakeProductRepository) FindByID(ctx context.Context, productID string) (*product.Product, error) {
	return f.products[productID], nil
}

func TestService_Create(t *testing.T) {
	repository := &fakeRepository{}

	productRepository := &fakeProductRepository{
		products: map[string]*product.Product{
			"product-1": {
				ID:    "product-1",
				Name:  "keyboard",
				Price: 59,
			},
			"product-2": {
				ID:    "product-2",
				Name:  "mouse",
				Price: 25,
			},
		},
	}

	service := order.NewService(repository, productRepository)

	req := order.CreateOrderRequest{
		Items: []order.CreateOrderItemRequest{
			{
				ProductID: "product-1",
				Quantity:  2,
			},
			{
				ProductID: "product-2",
				Quantity:  1,
			},
		},
	}

	order_test, err := service.Create(
		context.Background(), "user-1", req,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if order_test == nil {
		t.Fatalf("expected order, got nil")
	}

	expectedTotal := 143.0

	if order_test.TotalAmount != expectedTotal {
		t.Fatalf(
			"expected total %.2f, got %.2f", expectedTotal, order_test.TotalAmount,
		)
	}

	if order_test.Status != order.StatusPending {
		t.Fatalf("expected status %q, got %q", order.StatusPending, order_test.Status)
	}

	if len(repository.createdItems) != 2 {
		t.Fatalf("expected 2 order items, got %d", len(repository.createdItems))
	}

	first := repository.createdItems[0]

	if first.ProductName != "keyboard" {
		t.Fatalf("expected product name %q, got %q", "keyboard", first.ProductName)
	}

	if first.UnitPrice != 59 {
		t.Fatalf("expected unit price %d, got %.2f", 50, first.UnitPrice)
	}

	if first.Quantity != 2 {
		t.Fatalf("expected quantity %d,got %d", 2, first.Quantity)
	}

	if first.Subtotal != 118 {
		t.Fatalf(
			"expected subtotal %d,got %.2f", 118, first.Subtotal,
		)
	}
}
