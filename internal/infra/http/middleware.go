package http

import "net/http"

func JwtValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement JWT middleware logic here
		next.ServeHTTP(w, r)
	})
}
