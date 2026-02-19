package auth

import (
	"net/http"
	"strings"

	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type BasicAuth struct {
	SecretKey string
}

func NewBasicAuth(secretKey string) *BasicAuth {
	return &BasicAuth{SecretKey: secretKey}
}

func (a *BasicAuth) ValidateToken(request *http.Request, tokenType, tokenString string) (*http.Request, error) {
	if strings.ToLower(tokenType) != "basic" {
		return nil, pedalea.ErrInvalidTokenFormat
	}

	if a.SecretKey != tokenString {
		return nil, pedalea.ErrInvalidTokenCredentials
	}

	return request, nil
}
