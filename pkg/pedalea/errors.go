package pedalea

import "errors"

var (
	ErrInvalidToken = errors.New("Invalid token")

	ErrInvalidJwtSubject  = errors.New("Jwt subject is invalid id")
	ErrInvalidCredentials = errors.New("Invalid credentials")

	ErrInvalidOperation = errors.New("Invalid operation")

	ErrBikeAlreadyRented = errors.New("Bike is already rented")
	ErrUserAlreadyRented = errors.New("User already rented a bike")
)
