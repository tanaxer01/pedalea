package pedalea

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("User already exists")
	ErrInvalidJwtSubject  = errors.New("Jwt subject is invalid id")
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrUserNotFound       = errors.New("User not found")
)
