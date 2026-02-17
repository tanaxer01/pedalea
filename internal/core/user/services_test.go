package user

//go:generate mockery --name=UserRepo --output=./mocks --outpkg=mocks
//go:generate mockery --name=Crypto --output=./mocks --outpkg=mocks
//go:generate mockery --name=Auth --output=./mocks --outpkg=mocks

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/tanaxer01/pedalea/internal/core/user/mocks"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

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

	repo.On("GetUserByEmail", mock.Anything).Return(&pedalea.User{HashedPassword: "hashed"}, nil)
	crypto.On("ValidatePassword", "hashed", "pasword").Return(pedalea.ErrInvalidCredentials)

	s := NewService(repo, crypto, nil)

	token, err := s.Login(pedalea.LoginUser{
		Email:    "test@test.com",
		Password: "pasword",
	})

	require.Empty(t, token)
	require.ErrorIs(t, err, pedalea.ErrInvalidCredentials)

	repo.AssertExpectations(t)
	crypto.AssertExpectations(t)
}

func TestGetNonExistingUser(t *testing.T) {
	repo := new(mocks.UserRepo)

	repo.On("GetUserByID", mock.Anything).Return((*pedalea.User)(nil), pedalea.ErrUserNotFound)

	s := NewService(repo, nil, nil)

	user, err := s.GetUserData(1)

	require.ErrorIs(t, err, pedalea.ErrUserNotFound)
	require.Empty(t, user)

	repo.AssertExpectations(t)
}
