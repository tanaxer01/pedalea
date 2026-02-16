package rental

import (
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
	ListRentals(userID int) ([]pedalea.Rental, error)
	ListOverlappingRentals(userID, bikeID int) ([]pedalea.Rental, error)
}

func NewService(bikeRepo BikeRepo, rentalRepo RentalRepo) *Service {
	return &Service{bikeRepo: bikeRepo, rentalRepo: rentalRepo}
}

func (s *Service) StartRental(ID int, event pedalea.RentalEvent) error {
	// We ensure users can't start rentals for other users
	if ID != event.UserID {
		return pedalea.ErrInvalidOperation
	}

	rentals, err := s.rentalRepo.ListOverlappingRentals(event.UserID, event.BikeID)
	if err != nil {
		return err
	}

	// We ensure only valid users & bikes can be used to start a rental
	for _, rental := range rentals {
		if rental.UserID == event.UserID {
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

	// TODO: Do they want to define the end coords here ???
	err = s.rentalRepo.InsertRental(pedalea.RentalData{
		UserID:         event.UserID,
		BikeID:         event.BikeID,
		StartTime:      time.Now(),
		StartLatitude:  bike.Latitude,
		StartLongitude: bike.Longitude,
	})

	return err
}

func (s *Service) EndRental(UserID int) error {
	// TODO: Having partial updates would make this much easier
	// err := s.rentalRepo.UpdateRental(id, )

	// err := s.bikeRepo.UpdateBike()

	return nil
}

func (s *Service) ListRentals(ID int) ([]pedalea.RentalData, error) {
	rentals, err := s.rentalRepo.ListRentals(ID)
	if err != nil {
		return nil, err
	}

	data := make([]pedalea.RentalData, 0, len(rentals))
	for _, rental := range rentals {
		data = append(data, rental.RentalData)
	}

	return data, nil
}
