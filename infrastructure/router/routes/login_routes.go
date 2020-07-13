package routes

import (
	"merpochi_server/infrastructure/auth"
	"merpochi_server/infrastructure/database"
	"merpochi_server/interfaces/handler"
	"merpochi_server/usecase"
	"net/http"
)

func iniLoginRoutes() []Route {
	// 依存関係を注入
	authPersistence := auth.NewAuthPersistence(database.DB)
	authUsecase := usecase.NewAuthUsecase(authPersistence)
	authHandler := handler.NewAuthHandler(authUsecase)

	loginRoutes := []Route{
		{
			URI:          "/login",
			Method:       http.MethodPost,
			Handler:      authHandler.HandleLogin,
			AuthRequired: false,
		},
		{
			URI:          "/verify",
			Method:       http.MethodGet,
			Handler:      authHandler.HandleVerify,
			AuthRequired: true,
		},
	}
	return loginRoutes
}
