package order

import "errors"

var (
	ErrEmptyOrders     = errors.New("order must contain at least one item")
	ErrInvalidQuantity = errors.New("quanity must be greater than zero")
	ErrProductNotFound = errors.New("product not found")
)
