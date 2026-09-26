package payment

import (
	"context"

	"github.com/google/uuid"
)

var _ Provider = (*FakeProvider)(nil)

type FakeProvider struct {
	ShouldFail bool
}

func NewFakeProvider() *FakeProvider {
	return &FakeProvider{}
}

func (p *FakeProvider) Charge(ctx context.Context, amount float64) (string, error) {
	if p.ShouldFail {
		return "", ErrPaymentFailed
	}
	return uuid.NewString(), nil
}
