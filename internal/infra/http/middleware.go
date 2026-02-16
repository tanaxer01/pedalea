package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/tanaxer01/pedalea/internal/infra/auth"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
	"github.com/tanaxer01/pedalea/pkg/utils"
)

type JwtMiddleware struct {
	auth *auth.Auth
}

func NewJwtMiddleware(auth *auth.Auth) JwtMiddleware {
	return JwtMiddleware{auth: auth}
}

func (m JwtMiddleware) JwtValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, pedalea.ErrInvalidToken)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		token = strings.TrimPrefix(token, "bearer ")

		if token == "" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, pedalea.ErrInvalidToken)
			return
		}

		claim, err := m.auth.ValidateJwtToken(token)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, err)
			return
		}

		subject, err := claim.GetSubject()
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, err)
			return
		}

		ctx := context.WithValue(r.Context(), "UserID", subject)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
