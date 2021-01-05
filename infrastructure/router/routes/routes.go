package routes

import (
	"merpochi_server/config"
	"merpochi_server/infrastructure/middlewares"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Route ルーティング値の保管
type Route struct {
	URI          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Load ルーティング値の読込
func Load() []Route {
	routes := iniUsersRoutes()
	routes = append(routes, iniAuthRoutes()...)
	routes = append(routes, iniShopsRoutes()...)
	routes = append(routes, iniFavoritesRoutes()...)
	routes = append(routes, iniCommentsRoutes()...)
	routes = append(routes, iniStationsRoutes()...)
	routes = append(routes, iniImagesRoutes()...)
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
		if route.AuthRequired {
			r.HandleFunc(route.URI,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						middlewares.SetMiddlewareAuthentication(route.Handler))),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(route.Handler)),
			).Methods(route.Method)
		}
	}
	r.Use(cors)
	return r
}
