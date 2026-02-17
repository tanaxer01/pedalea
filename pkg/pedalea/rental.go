package pedalea

import (
	"errors"
)

type RentalStatus string

var (
	ErrRentalAlreadyExists = errors.New("rental already exists")
	ErrRentalNotFound      = errors.New("Rental not found")

	ErrBikeAlreadyRented = errors.New("Bike is already rented")
	ErrUserAlreadyRented = errors.New("User already rented a bike")

	ErrRentalInvalidEndCoords = errors.New("End coordinates are outside valid range")
)

const (
	StatusRunning RentalStatus = "running"
	StatusStopped RentalStatus = "ended"
)

type RentalData struct {
	UserID         int          `json:"user_id" db:"user_id"`
	BikeID         int          `json:"bike_id" db:"bike_id"`
	Status         RentalStatus `json:"status" db:"status"`
	StartTime      int64        `json:"start_time" db:"start_time"`
	EndTime        *int64       `json:"end_time" db:"end_time"`
	StartLatitude  float64      `json:"start_latitude" db:"start_latitude"`
	StartLongitude float64      `json:"start_longitude" db:"start_longitude"`
	EndLatitude    *float64     `json:"end_latitude" db:"end_latitude"`
	EndLongitude   *float64     `json:"end_longitude" db:"end_longitude"`
}

type StartRental struct {
	BikeID int
}

type EndRental struct {
	EndLatitude  float64
	EndLongitude float64
}

type Rental struct {
	ID       int
	CreateAt string
	UpdateAt string
	RentalData
}
