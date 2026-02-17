package rental

import (
	"math"
	"time"

	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type Service struct {
	bikeRepo   BikeRepo
	rentalRepo RentalRepo
}

type BikeRepo interface {
	UpdateBike(bikeID int, data pedalea.BikeData) error
	GetBikeByID(bikeID int) (*pedalea.Bike, error)
}

type RentalRepo interface {
	InsertRental(data pedalea.RentalData) error
	UpdateRental(userID int, data pedalea.RentalData) error
	GetRentalByUserID(ID int) (*pedalea.Rental, error)
	ListUserRentals(userID int) ([]pedalea.Rental, error)
	ListOverlappingRentals(userID, bikeID int) ([]pedalea.Rental, error)
}

func NewService(bikeRepo BikeRepo, rentalRepo RentalRepo) *Service {
	return &Service{bikeRepo: bikeRepo, rentalRepo: rentalRepo}
}

func (s *Service) StartRental(ID int, event pedalea.StartRental) error {
	// We ensure users can't start rentals for other users
	rentals, err := s.rentalRepo.ListOverlappingRentals(ID, event.BikeID)
	if err != nil {
		return err
	}

	// We ensure only valid users & bikes can be used to start a rental
	for _, rental := range rentals {
		if rental.UserID == ID {
			return pedalea.ErrUserAlreadyRented
		}

		if rental.BikeID == event.BikeID {
			return pedalea.ErrBikeAlreadyRented
		}
	}

	bike, err := s.bikeRepo.GetBikeByID(event.BikeID)
	if err != nil {
		return err
	}

	// We update the availability of the bike
	err = s.bikeRepo.UpdateBike(event.BikeID, pedalea.BikeData{
		Available: false,
		Latitude:  bike.Latitude,
		Longitude: bike.Longitude,
	})
	if err != nil {
		return nil
	}

	err = s.rentalRepo.InsertRental(pedalea.RentalData{
		UserID:         ID,
		BikeID:         event.BikeID,
		StartLatitude:  bike.Latitude,
		StartLongitude: bike.Longitude,
	})

	return err
}

func (s *Service) EndRental(UserID int, event pedalea.EndRental) error {
	rental, err := s.rentalRepo.GetRentalByUserID(UserID)
	if err != nil {
		return err
	}

	// Haversine formula
	// https://stackoverflow.com/questions/4913349/haversine-formula-in-python-bearing-and-distance-between-two-gps-points
	R := 6371. // earth radius in km

	toRad := func(d float64) float64 { return d * math.Pi / 180 }

	lat1 := toRad(rental.StartLatitude)
	lat2 := toRad(event.EndLatitude)
	dlat := toRad(event.EndLatitude - rental.StartLatitude)
	dlon := toRad(event.EndLongitude - rental.StartLongitude)

	a := math.Pow(math.Sin(dlat/2), 2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	if R*c > 5 {
		return pedalea.ErrRentalInvalidEndCoords
	}

	err = s.bikeRepo.UpdateBike(rental.BikeID, pedalea.BikeData{
		Available: true,
		Latitude:  event.EndLatitude,
		Longitude: event.EndLongitude,
	})

	if err != nil {
		return err
	}

	err = s.rentalRepo.UpdateRental(rental.ID, pedalea.RentalData{
		Status:       pedalea.StatusStopped,
		EndTime:      time.Now(),
		EndLatitude:  event.EndLatitude,
		EndLongitude: event.EndLongitude,
	})

	return err
}

func (s *Service) ListRentals(ID int) ([]pedalea.RentalData, error) {
	rentals, err := s.rentalRepo.ListUserRentals(ID)
	if err != nil {
		return nil, err
	}

	data := make([]pedalea.RentalData, 0, len(rentals))
	for _, rental := range rentals {
		data = append(data, rental.RentalData)
	}

	return data, nil
}
