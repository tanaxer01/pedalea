package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type Auth struct {
	SecretKey string
}

func NewAuth(secretKey string) *Auth {
	return &Auth{SecretKey: secretKey}
}

func (a *Auth) GenerateJwtToken(claims pedalea.UserClaim) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *Auth) ValidateJwtToken(tokenString string) (*pedalea.UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &pedalea.UserClaim{}, func(token *jwt.Token) (any, error) {
		return []byte(a.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*pedalea.UserClaim)
	if !ok || !token.Valid {
		return nil, pedalea.ErrInvalidToken
	}

	return claims, nil
}
