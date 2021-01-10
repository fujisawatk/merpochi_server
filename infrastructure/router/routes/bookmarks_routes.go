package routes

import (
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/persistence"
	"merpochi_server/interfaces/handler"
	"merpochi_server/usecase"
	"net/http"
)

func iniBookmarksRoutes() []Route {
	// 依存関係を注入
	bookmarkPersistence := persistence.NewBookmarkPersistence(database.DB)
	bookmarkUsecase := usecase.NewBookmarkUsecase(bookmarkPersistence)
	bookmarkHandler := handler.NewBookmarkHandler(bookmarkUsecase)

	bookmarksRoutes := []Route{
		{
			URI:          "/shops/{shopId}/bookmarks",
			Method:       http.MethodPost,
			Handler:      bookmarkHandler.HandleBookmarkCreate,
			AuthRequired: true,
		},
	}
	return bookmarksRoutes
}
