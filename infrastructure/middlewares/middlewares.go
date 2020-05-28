package middlewares

import (
	"log"
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

// SetMiddlewareLogger ルーティングのログ表示
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 形式を設定してログ表示（例：GET localhost:9000/users HTTP/1.1）
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

// SetMiddlewareJSON レスポンスをjson形式で返却
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
