package crypto

import (
	"errors"

	"github.com/tanaxer01/pedalea/pkg/pedalea"
	"golang.org/x/crypto/bcrypt"
)

type Crypto struct{}

func (c *Crypto) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (c *Crypto) ValidatePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return pedalea.ErrInvalidCredentials
	}

	return err
}
