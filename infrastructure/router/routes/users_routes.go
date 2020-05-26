package routes

import (
	"merpochi_server/interfaces/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	{
		URI:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
	{
		URI:     "/users/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	{
		URI:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
}
