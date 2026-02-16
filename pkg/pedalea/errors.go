package pedalea

import "errors"

var (
	ErrInvalidToken = errors.New("Invalid token")

	ErrUserAlreadyExists  = errors.New("User already exists")
	ErrInvalidJwtSubject  = errors.New("Jwt subject is invalid id")
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrUserNotFound       = errors.New("User not found")
)
