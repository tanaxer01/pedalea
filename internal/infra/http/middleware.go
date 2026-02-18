package http

import (
	"net/http"
	"strings"

	"github.com/tanaxer01/pedalea/pkg/pedalea"
	"github.com/tanaxer01/pedalea/pkg/utils"
)

type Auth interface {
	ValidateToken(request *http.Request, tokenType, tokenString string) (*http.Request, error)
}

type AuthMiddleware struct {
	auth Auth
}

func NewAuthMiddleware(auth Auth) *AuthMiddleware {
	return &AuthMiddleware{auth: auth}
}

func (m *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, pedalea.ErrInvalidTokenCredentials)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, pedalea.ErrInvalidTokenFormat)
			return
		}

		r, err := m.auth.ValidateToken(r, tokenParts[0], tokenParts[1])
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
