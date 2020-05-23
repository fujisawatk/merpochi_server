package routes

import (
	"merpochi_server/api/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
}
