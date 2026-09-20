package payment

import "context"

type Repository interface{
	Create(ctx context.Context, payment *Payment)error

	UpdateStatus(ctx context.Context, paymentID string, status Status, providerRef *string,)error
}


