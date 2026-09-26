package payment

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/0xlakhe/Impluse/internal/database"
	"github.com/google/uuid"
)

type Service struct {
	repository Repository
	provider   Provider
}

func NewService(repository Repository, provider Provider) *Service {
	return &Service{
		repository: repository, provider: provider,
	}
}

func (s *Service) Charge(ctx context.Context, orderId string, amount float64, idempotencyKey string) (*Payment, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	if idempotencyKey == "" {
		return nil, errors.New("idempotency key is required")
	}

	existing, err := s.repository.FindByIdempotencyKey(
		ctx, idempotencyKey,
	)

	if err == nil {
		return existing, nil
	}

	if !database.IsNotFound(err) {
		return nil, err
	}

	now := time.Now()

	payment := &Payment{
		ID:             uuid.NewString(),
		OrderID:        orderId,
		Amount:         amount,
		Status:         StatusPending,
		IdempotencyKey: idempotencyKey,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.repository.Create(ctx, payment); err != nil {
		return nil, fmt.Errorf("create payment: %w", err)
	}

	providerRef, err := s.provider.Charge(ctx, amount)
	if err != nil {
		_ = s.repository.UpdateStatus(ctx, payment.ID, StatusFailed, nil)
		payment.Status = StatusFailed
		return payment, err
	}

	if err := s.repository.UpdateStatus(ctx, payment.ID, StatusSuccess, &providerRef); err != nil {
		return nil, fmt.Errorf("update successful payment:  %w", err)
	}

	payment.Status = StatusSuccess
	payment.ProviderRef = &providerRef

	return payment, nil

}
