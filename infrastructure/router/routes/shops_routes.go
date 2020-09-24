package routes

import (
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/persistence"
	"merpochi_server/interfaces/handler"
	"merpochi_server/usecase"
	"net/http"
)

func iniShopsRoutes() []Route {
	// 依存関係を注入
	shopPersistence := persistence.NewShopPersistence(database.DB)
	shopUsecase := usecase.NewShopUsecase(shopPersistence)
	shopHandler := handler.NewShopHandler(shopUsecase)

	shopsRoutes := []Route{
		{
			URI:          "/shops",
			Method:       http.MethodPost,
			Handler:      shopHandler.HandleShopsGet,
			AuthRequired: false,
		},
		{
			URI:          "/shops/register",
			Method:       http.MethodPost,
			Handler:      shopHandler.HandleShopCreate,
			AuthRequired: false,
		},
	}
	return shopsRoutes
}
