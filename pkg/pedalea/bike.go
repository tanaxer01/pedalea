package pedalea

type BikeData struct {
	Available bool
	Latitude  bool
	Longitude bool
}

type Bike struct {
	ID       int
	CreateAt string
	UpdateAt string
	BikeData
}
