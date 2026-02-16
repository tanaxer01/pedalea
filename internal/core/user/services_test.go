package user

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type MockUserRepo struct{ mock.Mock }
type MockCrypto struct{ mock.Mock }
type MockAuth struct{ mock.Mock }

// TODO: This is awfull, look for a lib for mocking
func (m *MockUserRepo) InsertUser(data *pedalea.InsertUser) error {
	args := m.Called(data)
	return args.Error(0)
}
func (m *MockUserRepo) UpdateUser(ID int, data pedalea.UserData) error {
	args := m.Called(ID, data)
	return args.Error(0)
}
func (m *MockUserRepo) GetUserByID(ID int) (*pedalea.User, error) {
	args := m.Called(ID)
	user, _ := args.Get(0).(*pedalea.User)
	return user, args.Error(1)
}
func (m *MockUserRepo) GetUserByEmail(email string) (*pedalea.User, error) {
	args := m.Called(email)
	user, _ := args.Get(0).(*pedalea.User)
	return user, args.Error(1)
}

func (m *MockCrypto) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockCrypto) ValidatePassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(1)
}

func (m *MockAuth) GenerateJwtToken(params map[string]any) (string, error) {
	args := m.Called(params)
	return args.String(0), args.Error(1)
}

func TestInsertExistingUser(t *testing.T) {
	repo := new(MockUserRepo)
	crypto := new(MockCrypto)

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

// TestLoginNonexistingUser()

func TestLoginWrongPassword(t *testing.T) {
	repo := new(MockUserRepo)
	crypto := new(MockCrypto)
	auth := new(MockAuth)

	repo.On("GetUserByEmail", mock.Anything).Return(&pedalea.User{}, nil)
	crypto.On("ValidatePassword", mock.Anything, mock.Anything).Return(nil, pedalea.ErrInvalidCredentials)

	s := NewService(repo, crypto, auth)

	token, err := s.Login(pedalea.LoginUser{
		Email:    "test@test.com",
		Password: "pasword",
	})

	require.Empty(t, token)
	require.ErrorIs(t, err, pedalea.ErrInvalidCredentials)

	repo.AssertExpectations(t)
	crypto.AssertExpectations(t)
	auth.AssertExpectations(t)
}

// func TestGetUserWithInvalidJwt(t *testing.T) {}

// func TestGetUserWithMissingUser(t *testing.T) {}
