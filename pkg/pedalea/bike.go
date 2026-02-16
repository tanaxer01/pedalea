package pedalea

type BikeData struct {
	Available bool
	Latitude  float64
	Longitude float64
}

type Bike struct {
	ID        int
	CreatedAt string
	UpdatedAt string
	BikeData
}
