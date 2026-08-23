package user

import "errors"

var(
	ErrEmailAlreadyExists=errors.New("email already registered")
	ErrUserNameTaken=errors.New("username already taken")
	ErrInvalidCredentials=errors.New("Invalid credentials")
	ErrUserNotFound=errors.New("user not found")
)