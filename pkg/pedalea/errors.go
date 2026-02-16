package pedalea

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidJwtSubject  = errors.New("Jwt subject is invalid id")
	ErrInvalidCredentials = errors.New("Invalid credentials")
)
