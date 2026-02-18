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
		Latitude:  0,
		Longitude: 0,
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

	err = rentalService.EndRental(1, pedalea.EndRental{EndLatitude: 0, EndLongitude: 0})
	require.Nil(t, err)

	rental, err := rentalRepo.GetRentalByID(1)
	require.Nil(t, err)
	require.Equal(t, pedalea.StatusStopped, rental.Status)

	rentals, err := rentalRepo.ListRentals()
	require.Nil(t, err)

	require.Contains(t, rentals, *rental)
}
