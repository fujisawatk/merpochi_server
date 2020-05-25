package routes

import (
	"merpochi_server/config"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Route ルーティング値の保管
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// Load ルーティング値の読込
func Load() []Route {
	routes := usersRoutes
	return routes
}

// SetupRoutes ルーティングの設定
func SetupRoutes(r *mux.Router) *mux.Router {
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{config.CLIENTURL}),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"}),
		handlers.AllowedHeaders([]string{"Authorization"}),
	)
	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	r.Use(cors)
	return r
}
