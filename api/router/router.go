package router

import (
	"merpochi_server/api/router/routes"

	"github.com/gorilla/mux"
)

// New ルーティングの宣言
func New() *mux.Router {
	r := mux.NewRouter()
	return routes.SetupRoutes(r)
}
