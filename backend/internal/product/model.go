package product

import "time"

type Product struct {
	ID          string
	SellerID    string
	Name        string
	Description string
	Price       float64
	ImageURL    *string
	Category    string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
