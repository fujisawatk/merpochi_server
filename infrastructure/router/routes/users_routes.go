package routes

import (
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/persistence"
	"merpochi_server/interfaces/handler"
	"merpochi_server/usecase"
	"net/http"
)

func iniUserRoutes() []Route {
	// 依存関係を注入
	userPersistence := persistence.NewUserPersistence(database.DB)
	userUsecase := usecase.NewUserUsecase(userPersistence)
	userHandler := handler.NewUserHandler(userUsecase)

	usersRoutes := []Route{
		{
			URI:          "/users",
			Method:       http.MethodGet,
			Handler:      userHandler.HandleUsersGet,
			AuthRequired: false,
		},
		// {
		// 	URI:          "/users",
		// 	Method:       http.MethodPost,
		// 	Handler:      controllers.CreateUser,
		// 	AuthRequired: true,
		// },
		// {
		// 	URI:          "/users/{id}",
		// 	Method:       http.MethodGet,
		// 	Handler:      controllers.GetUser,
		// 	AuthRequired: false,
		// },
		// {
		// 	URI:          "/users/{id}",
		// 	Method:       http.MethodPut,
		// 	Handler:      controllers.UpdateUser,
		// 	AuthRequired: true,
		// },
		// {
		// 	URI:          "/users/{id}",
		// 	Method:       http.MethodDelete,
		// 	Handler:      controllers.DeleteUser,
		// 	AuthRequired: true,
		// },
	}
	return usersRoutes
}
