package payment

import "time"

type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusSuccess    Status = "success"
	StatusFailed     Status = "failed"
)

type Payment struct {
	ID             string
	OrderID        string
	Amount         float64
	Status         Status
	ProviderRef    *string
	IdempotencyKey string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
