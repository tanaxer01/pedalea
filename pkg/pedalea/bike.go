package pedalea

import "errors"

var (
	ErrBikeAlreadyExists = errors.New("bike already exists")
	ErrBikeNotFound      = errors.New("bike not found")
)

type BikeData struct {
	Available bool    `json:"is_available" db:"is_available"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}

type Bike struct {
	ID        int    `db:"id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	BikeData
}
