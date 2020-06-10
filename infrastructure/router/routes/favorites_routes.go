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
			URI:          "/favorites/{id}",
			Method:       http.MethodPost,
			Handler:      favoriteHandler.HandleFavoriteCreate,
			AuthRequired: false,
		},
		{
			URI:          "/favorites/{id}",
			Method:       http.MethodDelete,
			Handler:      favoriteHandler.HandleFavoriteDelete,
			AuthRequired: false,
		},
	}
	return favoritesRoutes
}
