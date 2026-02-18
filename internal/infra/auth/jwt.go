package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type JwtAuth struct {
	SecretKey string
}

func NewJwtAuth(secretKey string) *JwtAuth {
	return &JwtAuth{SecretKey: secretKey}
}

func (a *JwtAuth) GenerateJwtToken(claims pedalea.UserClaim) (string, error) {
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

func (a *JwtAuth) ValidateToken(request *http.Request, tokenType, tokenString string) (*http.Request, error) {
	if strings.ToLower(tokenType) != "bearer" {
		return nil, pedalea.ErrInvalidTokenFormat
	}

	token, err := jwt.ParseWithClaims(tokenString, &pedalea.UserClaim{}, func(token *jwt.Token) (any, error) {
		return []byte(a.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*pedalea.UserClaim)
	if !ok || !token.Valid {
		return nil, pedalea.ErrInvalidTokenCredentials
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return nil, pedalea.ErrInvalidTokenFormat
	}

	ctx := context.WithValue(request.Context(), "UserID", subject)
	return request.WithContext(ctx), nil
}
