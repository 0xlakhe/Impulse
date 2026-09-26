package order

import "context"

type Repository interface{
	Create(ctx context.Context, order *Order, items []OrderItem,)error

	UpdateStatus(ctx context.Context, orderID string, status Status,)error
}

