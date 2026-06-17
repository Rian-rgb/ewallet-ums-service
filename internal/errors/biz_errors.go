package errors

import "errors"

var (
	ErrUserNotFound        = errors.New("user account not found")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrInvalidToken        = errors.New("invalid token")
	ErrInternalServerError = errors.New("an unexpected errors occurred. Please try again later")
)
