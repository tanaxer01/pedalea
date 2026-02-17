package pedalea

import (
	"errors"
	"time"
)

type RentalStatus string

var (
	ErrRentalAlreadyExists = errors.New("rental already exists")
	ErrRentalNotFound      = errors.New("Rental not found")

	ErrRentalInvalidEndCoords = errors.New("End coordinates are outside valid range")
)

const (
	StatusRunning RentalStatus = "running"
	StatusStopped RentalStatus = "ended"
)

type RentalData struct {
	UserID         int
	BikeID         int
	Status         RentalStatus
	StartTime      time.Time
	EndTime        time.Time
	StartLatitude  float64
	StartLongitude float64
	EndLatitude    float64
	EndLongitude   float64
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
