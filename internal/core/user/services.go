package user

import (
	"time"

	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type Service struct {
	userRepo UserRepo
	crypto   Crypto
	auth     Auth
}

type UserRepo interface {
	InsertUser(data *pedalea.InsertUser) error
	UpdateUser(ID int, data pedalea.UserData) error
	GetUserByID(ID int) (*pedalea.User, error)
	GetUserByEmail(email string) (*pedalea.User, error)
}

type Crypto interface {
	HashPassword(password string) (string, error)
	ValidatePassword(hashedPassword, password string) error
}

type Auth interface {
	GenerateJwtToken(params map[string]any) (string, error)
}

func NewService(userRepo UserRepo, crypto Crypto, auth Auth) *Service {
	return &Service{userRepo: userRepo, crypto: crypto, auth: auth}
}

// TODO: Should we return a JWT token?
func (s *Service) InsertUser(data pedalea.InsertUser) error {
	hashedPassword, err := s.crypto.HashPassword(data.Password)
	if err != nil {
		return err
	}

	data.Password = hashedPassword
	err = s.userRepo.InsertUser(&data)

	return err
}

// TODO: This does not allow partial updates, that sucks
func (s *Service) UpdateUser(ID int, data pedalea.UserData) error {
	return s.userRepo.UpdateUser(ID, data)
}

func (s *Service) Login(data pedalea.LoginUser) (string, error) {
	user, err := s.userRepo.GetUserByEmail(data.Email)
	if err != nil {
		return "", err
	}

	if err := s.crypto.ValidatePassword(user.HashedPassword, data.Password); err != nil {
		return "", err
	}

	token, err := s.auth.GenerateJwtToken(map[string]any{
		"sub":       user.ID,
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GetUserData(ID int) (*pedalea.UserData, error) {
	user, err := s.userRepo.GetUserByID(ID)
	if err != nil {
		return nil, err
	}

	return &user.UserData, nil
}
