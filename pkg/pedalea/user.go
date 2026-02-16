package pedalea

import "github.com/golang-jwt/jwt/v5"

type UserData struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
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
	ID             int
	CreatedAt      string
	UpdatedAt      string
	HashedPassword string
	UserData
}

type UserClaim struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.RegisteredClaims
}
