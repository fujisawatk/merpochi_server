package routes

import (
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/persistence"
	"merpochi_server/interfaces/handler"
	"merpochi_server/usecase"
	"net/http"
)

func iniUsersRoutes() []Route {
	// 依存関係を注入
	userPersistence := persistence.NewUserPersistence(database.DB)
	userUsecase := usecase.NewUserUsecase(userPersistence)
	userHandler := handler.NewUserHandler(userUsecase)

	usersRoutes := []Route{
		{
			URI:          "/users",
			Method:       http.MethodPost,
			Handler:      userHandler.HandleUserCreate,
			AuthRequired: false,
		},
		{
			URI:          "/users/{id}",
			Method:       http.MethodGet,
			Handler:      userHandler.HandleUserGet,
			AuthRequired: false,
		},
		{
			URI:          "/users/{id}",
			Method:       http.MethodPut,
			Handler:      userHandler.HandleUserUpdate,
			AuthRequired: true,
		},
		{
			URI:          "/users/{id}",
			Method:       http.MethodDelete,
			Handler:      userHandler.HandleUserDelete,
			AuthRequired: true,
		},
		{
			URI:          "/users/mylist",
			Method:       http.MethodPost,
			Handler:      userHandler.HandleUserMylist,
			AuthRequired: true,
		},
	}
	return usersRoutes
}
