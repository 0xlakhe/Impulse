package order

import "time"

type Status string

const (
	StatusPending   Status = "pending"
	StatusPaid      Status = "paid"
	StatusFailed    Status = "failed"
	StatusCancelled Status = "cancelled"
)

type Order struct {
	ID          string
	UserID      string
	TotalAmount float64
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderItem struct {
	ID          string
	OrderID     string
	ProductID   string
	ProductName string
	UnitPrice   float64
	Quantity    int
	Subtotal    float64
}
