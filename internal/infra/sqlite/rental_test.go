package sqlite

import (
	"testing"

	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/tanaxer01/pedalea/internal/core/rental"
	"github.com/tanaxer01/pedalea/internal/core/user"
	"github.com/tanaxer01/pedalea/internal/infra/auth"
	"github.com/tanaxer01/pedalea/internal/infra/crypto"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

func TestRentalFlow(t *testing.T) {
	db, err := NewDB(":memory:")
	require.Nil(t, err)

	err = goose.SetDialect("sqlite3")
	require.Nil(t, err)

	err = goose.UpByOne(db.DB, "../../../migrations")
	require.Nil(t, err)

	authRepo := auth.NewJwtAuth("secret-key")

	userRepo := NewUserRepository(db)
	userService := user.NewService(userRepo, &crypto.Crypto{}, authRepo)

	bikeRepo := NewBikeRepository(db)
	rentalRepo := NewRentalRepository(db)
	rentalService := rental.NewService(bikeRepo, rentalRepo)

	err = bikeRepo.InsertBike(pedalea.BikeData{
		Latitude:  10,
		Longitude: 20,
	})
	require.Nil(t, err)

	err = userService.InsertUser(pedalea.InsertUser{
		UserData: pedalea.UserData{
			Email:     "test@test.com",
			FirstName: "first",
			LastName:  "last",
		},
		Password: "password",
	})
	require.Nil(t, err)

	_, err = userService.Login(pedalea.LoginUser{
		Email:    "test@test.com",
		Password: "password",
	})
	require.Nil(t, err)

	err = rentalService.StartRental(1, pedalea.StartRental{BikeID: 1})
	require.Nil(t, err)

	running, err := rentalRepo.GetRentalByUserID(1)
	require.Nil(t, err)
	require.NotZero(t, running.StartTime)
	require.Equal(t, 10.0, running.StartLatitude)
	require.Equal(t, 20.0, running.StartLongitude)

	err = rentalService.EndRental(1, pedalea.EndRental{EndLatitude: 10, EndLongitude: 20})
	require.Nil(t, err)

	rental, err := rentalRepo.GetRentalByID(1)
	require.Nil(t, err)
	require.Equal(t, pedalea.StatusStopped, rental.Status)
	require.Equal(t, running.StartTime, rental.StartTime)
	require.Equal(t, running.StartLatitude, rental.StartLatitude)
	require.Equal(t, running.StartLongitude, rental.StartLongitude)

	rentals, err := rentalRepo.ListRentals()
	require.Nil(t, err)

	require.Contains(t, rentals, *rental)
}
