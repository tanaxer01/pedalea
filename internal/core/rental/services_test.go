package rental

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tanaxer01/pedalea/internal/core/rental/mocks"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

//go:generate mockery --name=BikeRepo --output=./mocks --outpkg=mocks
//go:generate mockery --name=RentalRepo --output=./mocks --outpkg=mocks

func TestRentalWithOverlappingUser(t *testing.T) {
	bikeRepo := new(mocks.BikeRepo)
	rentalRepo := new(mocks.RentalRepo)

	rentalRepo.On("ListOverlappingRentals", mock.Anything, mock.Anything).Return([]pedalea.Rental{
		{
			ID: 1,
			RentalData: pedalea.RentalData{
				UserID: 1,
				BikeID: 4,
			},
		},
	}, nil)

	err := NewService(bikeRepo, rentalRepo).StartRental(1, pedalea.StartRental{BikeID: 2})

	require.ErrorIs(t, err, pedalea.ErrUserAlreadyRented)

	bikeRepo.AssertNotCalled(t, "GetBikeByID", mock.Anything)
	bikeRepo.AssertNotCalled(t, "UpdateBike", mock.Anything, mock.Anything)
	bikeRepo.AssertExpectations(t)

	rentalRepo.AssertNotCalled(t, "InsertRental", mock.Anything)
	rentalRepo.AssertExpectations(t)
}

func TestRentalWithOverlappingBike(t *testing.T) {
	bikeRepo := new(mocks.BikeRepo)
	rentalRepo := new(mocks.RentalRepo)

	rentalRepo.On("ListOverlappingRentals", mock.Anything, mock.Anything).Return([]pedalea.Rental{
		{
			ID: 1,
			RentalData: pedalea.RentalData{
				UserID: 100,
				BikeID: 1,
			},
		},
	}, nil)

	err := NewService(bikeRepo, rentalRepo).StartRental(1, pedalea.StartRental{
		BikeID: 1,
	})
	require.ErrorIs(t, err, pedalea.ErrBikeAlreadyRented)

	bikeRepo.AssertNotCalled(t, "GetBikeByID", mock.Anything)
	bikeRepo.AssertNotCalled(t, "UpdateBike", mock.Anything, mock.Anything)
	bikeRepo.AssertExpectations(t)

	rentalRepo.AssertNotCalled(t, "InsertRental", mock.Anything)
	rentalRepo.AssertExpectations(t)
}

func TestClosingRentalOutsideRange(t *testing.T) {
	rentalRepo := new(mocks.RentalRepo)
	bikeRepo := new(mocks.BikeRepo)

	rentalRepo.On("GetRentalByUserID", mock.Anything).Return(&pedalea.Rental{
		RentalData: pedalea.RentalData{
			BikeID:         1,
			StartTime:      time.Now().Unix(),
			StartLatitude:  0,
			StartLongitude: 0,
		},
	}, nil)

	bikeRepo.On("GetBikeByID", mock.Anything).Return(&pedalea.Bike{
		BikeData: pedalea.BikeData{
			PricePerMinute: 10,
		},
	}, nil)

	s := NewService(bikeRepo, rentalRepo)
	err := s.EndRental(1, pedalea.EndRental{
		EndLatitude:  0.1,
		EndLongitude: 0.1,
	})

	require.ErrorIs(t, err, pedalea.ErrRentalInvalidEndCoords)

	rentalRepo.AssertExpectations(t)
	bikeRepo.AssertExpectations(t)
}
