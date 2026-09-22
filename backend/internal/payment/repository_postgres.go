package payment

import (
	"context"
	"fmt"

	"github.com/0xlakhe/Impluse/internal/database"
)

type postgresRepository struct {
	db *database.Database
}

func NewRepository(db *database.Database) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, payment *Payment) error {
	_, err := r.db.Pool.Exec(
		ctx,
		`INSERT INTO payments(
			id,
			order_id,
			amount,
			status,
			provider_ref,
			idempotency_key,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, payment.ID, payment.OrderID, payment.Amount, payment.Status, payment.ProviderRef, payment.IdempotencyKey, payment.CreatedAt, payment.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("create payment: %w", err)
	}
	return nil
}

func (r *postgresRepository) UpdateStatus(ctx context.Context, paymentID string, status Status, providerRef *string) error {
	result, err := r.db.Pool.Exec(
		ctx,
		`UPDATE payments
		SET
			status=$1,
			provider_ref=$2,
			updated_at=NOW()
		WHERE id=$3`, status, providerRef, paymentID,
	)
	if err != nil {
		return fmt.Errorf("update payment: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("payment %s not found", paymentID)
	}
	return nil
}

func (r *postgresRepository) FindByIdempotencyKey(ctx context.Context, idempotencyKey string) (*Payment, error) {
	return nil, nil
}
