package middlewares

import (
	"merpochi_server/infrastructure/auth"
	"merpochi_server/interfaces/responses"
	"net/http"
)

// SetMiddlewareAuthentication 認証トークンを検証
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
