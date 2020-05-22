package routes

import (
	"merpochi_server/api/controllers"
	"net/http"
)

var testsRoutes = []Route{
	{
		URI:     "/public",
		Method:  http.MethodGet,
		Handler: controllers.Public,
	},
}
