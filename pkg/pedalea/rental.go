package pedalea

import "time"

type RentalStatus string

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

type RentalEvent struct {
	UserID int
	BikeID int
}

type Rental struct {
	ID       int
	CreateAt string
	UpdateAt string
	RentalData
}
