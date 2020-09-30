package routes

import (
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/persistence"
	"merpochi_server/interfaces/handler"
	"merpochi_server/usecase"
	"net/http"
)

func iniFavoritesRoutes() []Route {
	// 依存関係を注入
	favoritePersistence := persistence.NewFavoritePersistence(database.DB)
	favoriteUsecase := usecase.NewFavoriteUsecase(favoritePersistence)
	favoriteHandler := handler.NewFavoriteHandler(favoriteUsecase)

	favoritesRoutes := []Route{
		{
			URI:          "/shops/{shopId}/favorites",
			Method:       http.MethodGet,
			Handler:      favoriteHandler.HandleFavoritesGet,
			AuthRequired: false,
		},
		{
			URI:          "/shops/{shopId}/favorites",
			Method:       http.MethodPost,
			Handler:      favoriteHandler.HandleFavoriteCreate,
			AuthRequired: true,
		},
		{
			URI:          "/shops/{shopId}/favorites",
			Method:       http.MethodDelete,
			Handler:      favoriteHandler.HandleFavoriteDelete,
			AuthRequired: true,
		},
	}
	return favoritesRoutes
}
