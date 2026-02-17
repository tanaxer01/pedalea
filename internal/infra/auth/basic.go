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

	credentials := strings.Split(tokenString, ":")
	if len(credentials) != 2 {
		return nil, pedalea.ErrInvalidTokenFormat
	}

	if credentials[0] != "admin" || credentials[1] != a.SecretKey {
		return nil, pedalea.ErrInvalidTokenCredentials
	}

	return request, nil
}
