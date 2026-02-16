package pedalea

import "errors"

var (
	ErrInvalidToken = errors.New("Invalid token")

	ErrUserAlreadyExists  = errors.New("User already exists")
	ErrInvalidJwtSubject  = errors.New("Jwt subject is invalid id")
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrUserNotFound       = errors.New("User not found")

	ErrInvalidOperation = errors.New("Invalid operation")

	ErrRentalNotFound      = errors.New("Rental not found")
	ErrRentalAlreadyExists = errors.New("Rental already exists")
	ErrBikeAlreadyRented   = errors.New("Bike is already rented")
	ErrUserAlreadyRented   = errors.New("User already rented a bike")
)
