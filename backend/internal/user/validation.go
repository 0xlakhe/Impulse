package user

import (
	"errors"
	"strings"
)

func ValidateRegisterRequest(req RegisterRequest) error {
	if len(strings.TrimSpace(req.Username)) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if len(req.Username) > 50 {
		return errors.New("username is too long")
	}
	if !strings.Contains(req.Email, "@") {
		return errors.New("invalid email")
	}
	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	if len(req.Password) > 72 {
		return errors.New("password is too long")
	}
	return nil
}
