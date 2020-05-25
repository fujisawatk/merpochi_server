package routes

import (
	"merpochi_server/interfaces/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
}
