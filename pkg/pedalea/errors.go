package pedalea

import "errors"

var (
	ErrInvalidToken = errors.New("Invalid token")

	ErrInvalidJwtSubject  = errors.New("Jwt subject is invalid id")
	ErrInvalidCredentials = errors.New("Invalid credentials")

	ErrInvalidOperation = errors.New("Invalid operation")
)
