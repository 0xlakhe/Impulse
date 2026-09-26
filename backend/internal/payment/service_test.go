package payment

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

type fakeRepository struct {
	payment     *Payment
	findErr     error
	createErr   error
	updateErr   error
	createCalls int
	updateCalls int
	providerRef *string
}

func (f *fakeRepository) Create(ctx context.Context, payment *Payment) error {
	f.createCalls++

	if f.createErr != nil {
		return f.createErr
	}
	f.payment = payment

	return nil
}

func (f *fakeRepository) FindByIdempotencyKey(ctx context.Context, idempotency string) (*Payment, error) {
	if f.findErr != nil {
		return nil, f.findErr
	}
	return f.payment, nil
}

func (f *fakeRepository) UpdateStatus(ctx context.Context, paymentID string, status Status, providerRef *string) error {
	f.updateCalls++

	if f.updateErr != nil {
		return f.updateErr
	}

	if f.payment != nil {
		f.payment.Status = status
		f.payment.ProviderRef = providerRef
	}

	f.providerRef = providerRef

	return nil
}

type fakeProvider struct {
	transactionID string
	err           error
	calls         int
	amount        float64
}

func (f *fakeProvider) Charge(ctx context.Context, amount float64) (string, error) {
	f.calls++
	f.amount = amount

	if f.err != nil {
		return "", f.err
	}

	return f.transactionID, nil
}

func TestService_Charge_Success(t *testing.T) {
	repository := &fakeRepository{
		findErr: errors.New("not found"),
	}

	provider := &fakeProvider{
		transactionID: uuid.NewString(),
	}

	service := NewService(repository, provider)

	payment, err := service.Charge(context.Background(), "order-1", 100, "checkout-123")

	if err != nil {
		t.Fatalf("expected no error, go %v", err)
	}

	if payment == nil {
		t.Fatalf("expected payment, got nil")
	}

	if payment.Status != StatusSuccess {
		t.Fatalf("expectd status %q, got %q", StatusSuccess, payment.Status)
	}

	if payment.Amount != 100 {
		t.Fatalf("expected amount 100, got %.2f", payment.Amount)
	}

	if provider.calls != 1 {
		t.Fatalf("expected provider to be called one, got %d", provider.calls)
	}

	if repository.createCalls != 1 {
		t.Fatalf("expected payment to be created once, got %d", repository.createCalls)
	}

	if repository.updateCalls != 1 {
		t.Fatalf("expected status to be updated once, got %d", repository.updateCalls)
	}

	if payment.ProviderRef == nil {
		t.Fatalf("expected provider reference")
	}
}
