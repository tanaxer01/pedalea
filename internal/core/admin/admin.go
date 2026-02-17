package admin

import "github.com/tanaxer01/pedalea/pkg/pedalea"

type Service struct {
	userRepo   UserRepo
	bikeRepo   BikeRepo
	rentalRepo RentalRepo
}

type UserRepo interface {
	UpdateUser(ID int, data pedalea.UserData) error
	GetUserByID(ID int) (*pedalea.User, error)
	ListUsers() ([]pedalea.User, error)
}

type BikeRepo interface {
	InsertBike(data pedalea.BikeData) error
	UpdateBike(ID int, data pedalea.BikeData) error
	ListBikes() ([]pedalea.Bike, error)
}

type RentalRepo interface {
	UpdateRental(ID int, data pedalea.RentalData) error
	GetRentalByID(ID int) (*pedalea.Rental, error)
	ListRentals() ([]pedalea.Rental, error)
}

func NewService(userRepo UserRepo, bikeRepo BikeRepo, rentalRepo RentalRepo) *Service {
	return &Service{
		userRepo:   userRepo,
		bikeRepo:   bikeRepo,
		rentalRepo: rentalRepo,
	}
}
