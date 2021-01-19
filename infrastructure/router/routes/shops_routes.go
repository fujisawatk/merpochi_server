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
	postPersistence := persistence.NewPostPersistence(database.DB)
	favoritePersistence := persistence.NewFavoritePersistence(database.DB)
	bookmarkPersistence := persistence.NewBookmarkPersistence(database.DB)
	shopUsecase := usecase.NewShopUsecase(
		shopPersistence,
		postPersistence,
		favoritePersistence,
		bookmarkPersistence,
	)
	shopHandler := handler.NewShopHandler(shopUsecase)

	shopsRoutes := []Route{
		{
			URI:          "/shops/search",
			Method:       http.MethodPost,
			Handler:      shopHandler.HandleShopsSearch,
			AuthRequired: false,
		},
		{
			URI:          "/shops",
			Method:       http.MethodPost,
			Handler:      shopHandler.HandleShopCreate,
			AuthRequired: false,
		},
		{
			URI:          "/shops/code",
			Method:       http.MethodPost,
			Handler:      shopHandler.HandleShopGet,
			AuthRequired: false,
		},
		{
			URI:          "/shops/posted",
			Method:       http.MethodPost,
			Handler:      shopHandler.HandleShopGetPosted,
			AuthRequired: false,
		},
	}
	return shopsRoutes
}
