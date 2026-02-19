package user

//go:generate mockery --name=UserRepo --output=./mocks --outpkg=mocks
//go:generate mockery --name=Crypto --output=./mocks --outpkg=mocks
//go:generate mockery --name=Auth --output=./mocks --outpkg=mocks

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/tanaxer01/pedalea/internal/core/user/mocks"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

// Insert
func TestInsertExistingUser(t *testing.T) {
	repo := new(mocks.UserRepo)
	crypto := new(mocks.Crypto)

	crypto.On("HashPassword", "password").Return("hashed", nil)
	repo.On("InsertUser", mock.Anything).Return(pedalea.ErrUserAlreadyExists)

	s := NewService(repo, crypto, nil)

	err := s.InsertUser(pedalea.InsertUser{
		UserData: pedalea.UserData{
			Email:     "test@example.com",
			FirstName: "John",
			LastName:  "Doe",
		},
		Password: "password",
	})

	require.ErrorIs(t, err, pedalea.ErrUserAlreadyExists)

	repo.AssertExpectations(t)
	crypto.AssertExpectations(t)
}

// Update
func TestUpdateToExistingEmail(t *testing.T) {
	repo := new(mocks.UserRepo)

	repo.On("UpdateUser", mock.Anything, mock.Anything).Return(pedalea.ErrEmailAlreadyExists)

	s := NewService(repo, nil, nil)

	err := s.UpdateUser(1, pedalea.UserData{
		Email:     "existing@email.com",
		FirstName: "New name",
	})

	require.ErrorIs(t, err, pedalea.ErrEmailAlreadyExists)

	repo.AssertExpectations(t)
}

// Login
func TestSuccessfulLogin(t *testing.T) {
	repo := new(mocks.UserRepo)
	crypto := new(mocks.Crypto)
	auth := new(mocks.Auth)

	user := pedalea.User{
		ID: 1010,
		UserData: pedalea.UserData{
			Email:     "test@email.com",
			FirstName: "First",
			LastName:  "Last",
		},
		HashedPassword: "hashed",
	}

	repo.On("GetUserByEmail", "test@email.com").Return(&user, nil)
	crypto.On("ValidatePassword", "hashed", "clear").Return(nil)
	auth.On("GenerateJwtToken", mock.MatchedBy(func(c pedalea.UserClaim) bool {
		if c.RegisteredClaims.Subject != strconv.Itoa(user.ID) {
			return false
		}

		if c.Email != user.Email || c.FirstName != user.FirstName || c.LastName != user.LastName {
			return false
		}

		return c.RegisteredClaims.ExpiresAt != nil
	})).Return("token", nil)

	token, err := NewService(repo, crypto, auth).Login(pedalea.LoginUser{
		Email:    "test@email.com",
		Password: "clear",
	})

	require.NoError(t, err)
	require.Equal(t, "token", token)

	repo.AssertExpectations(t)
	crypto.AssertExpectations(t)
	auth.AssertExpectations(t)
}

func TestLoginWithNonExistingUser(t *testing.T) {
	repo := new(mocks.UserRepo)

	repo.On("GetUserByEmail", mock.Anything).Return((*pedalea.User)(nil), pedalea.ErrUserNotFound)

	s := NewService(repo, nil, nil)

	token, err := s.Login(pedalea.LoginUser{
		Email:    "test@test.com",
		Password: "test-password",
	})

	require.ErrorIs(t, err, pedalea.ErrUserNotFound)
	require.Empty(t, token)

	repo.AssertExpectations(t)
}

func TestLoginWrongPassword(t *testing.T) {
	repo := new(mocks.UserRepo)
	crypto := new(mocks.Crypto)
	auth := new(mocks.Auth)

	repo.On("GetUserByEmail", "test@test.com").Return(&pedalea.User{HashedPassword: "hashed"}, nil)
	crypto.On("ValidatePassword", "hashed", "pasword").Return(pedalea.ErrInvalidCredentials)

	token, err := NewService(repo, crypto, auth).Login(pedalea.LoginUser{
		Email:    "test@test.com",
		Password: "pasword",
	})

	require.ErrorIs(t, err, pedalea.ErrInvalidCredentials)
	require.Empty(t, token)

	auth.AssertNotCalled(t, "GenerateJwtToken", mock.Anything)
	repo.AssertExpectations(t)
	crypto.AssertExpectations(t)
}

// Get
func TestGetNonExistingUser(t *testing.T) {
	repo := new(mocks.UserRepo)

	repo.On("GetUserByID", mock.Anything).Return((*pedalea.User)(nil), pedalea.ErrUserNotFound)

	s := NewService(repo, nil, nil)

	user, err := s.GetUserData(1)

	require.ErrorIs(t, err, pedalea.ErrUserNotFound)
	require.Empty(t, user)

	repo.AssertExpectations(t)
}
