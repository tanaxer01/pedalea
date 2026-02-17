package sqlite

import (
	"strconv"
	"testing"

	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/tanaxer01/pedalea/internal/core/user"
	"github.com/tanaxer01/pedalea/internal/infra/auth"
	"github.com/tanaxer01/pedalea/internal/infra/crypto"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

func TestNormalFlow(t *testing.T) {
	db, err := NewDB(":memory:")
	require.Nil(t, err)

	err = goose.SetDialect("sqlite3")
	require.Nil(t, err)

	err = goose.Up(db.DB, "../../../migrations")
	require.Nil(t, err)

	authRepo := auth.NewAuth("secret-key")

	r := NewUserRepository(db)
	s := user.NewService(r, &crypto.Crypto{}, authRepo)

	err = s.InsertUser(pedalea.InsertUser{
		UserData: pedalea.UserData{
			Email:     "test@test.com",
			FirstName: "first",
			LastName:  "last",
		},
		Password: "password",
	})
	require.Nil(t, err)

	token, err := s.Login(pedalea.LoginUser{
		Email:    "test@test.com",
		Password: "password",
	})
	require.Nil(t, err)

	claims, err := authRepo.ValidateJwtToken(token)
	require.Nil(t, err)

	ID, err := strconv.Atoi(claims.Subject)
	require.Nil(t, err)

	user, err := s.GetUserData(ID)
	require.Nil(t, err)
	require.Equal(t, &pedalea.UserData{
		Email:     "test@test.com",
		FirstName: "first",
		LastName:  "last",
	}, user)

	err = s.UpdateUser(ID, pedalea.UserData{
		Email:     "test@test.com",
		FirstName: "new",
		LastName:  "new",
	})
	require.Nil(t, err)

	user, err = s.GetUserData(ID)
	require.Nil(t, err)
	require.Equal(t, &pedalea.UserData{
		Email:     "test@test.com",
		FirstName: "new",
		LastName:  "new",
	}, user)
}
