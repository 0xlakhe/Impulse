package payment

import "context"

type Provider interface {
	Charge(ctx context.Context, amount float64) (string, error)
}
