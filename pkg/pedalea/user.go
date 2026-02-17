package pedalea

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("User not found")
)

type UserData struct {
	Email     string `json:"email" db:"email" validate:"required,email"`
	FirstName string `json:"first_name" db:"first_name" validate:"required"`
	LastName  string `json:"last_name" db:"last_name" validate:"required"`
}

type InsertUser struct {
	Password string `json:"password" validate:"required"`
	UserData
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID             int    `db:"id"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
	HashedPassword string `db:"hashed_password"`
	UserData
}

type UserClaim struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.RegisteredClaims
}
