package conversation

import "errors"

var(
	ErrNotFound=errors.New("not found")
	ErrUnauthorized=errors.New("user not authorized")
	ErrBadRequest=errors.New("invalid request body")
	ErrConversationNotFound=errors.New("conversation not found")
	ErrEmptyMessage=errors.New("message cannot be empty")
)