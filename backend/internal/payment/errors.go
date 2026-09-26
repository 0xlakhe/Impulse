package payment

import "errors"

var (
	ErrInvalidAmount = errors.New("payment amount must be greater than zero")
	ErrPaymentFailed = errors.New("payment failed")
)
