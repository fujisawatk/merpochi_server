package routes

import (
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/persistence"
	"merpochi_server/interfaces/handler"
	"merpochi_server/usecase"
	"net/http"
)

func iniPostsRoutes() []Route {
	// 依存関係を注入
	postPersistence := persistence.NewPostPersistence(database.DB)
	commentPersistence := persistence.NewCommentPersistence(database.DB)
	imagePersistence := persistence.NewImagePersistence(database.DB)
	postUsecase := usecase.NewPostUsecase(postPersistence, commentPersistence, imagePersistence)
	postHandler := handler.NewPostHandler(postUsecase)

	postRoutes := []Route{
		{
			URI:          "/shops/{id}/posts",
			Method:       http.MethodPost,
			Handler:      postHandler.HandlePostCreate,
			AuthRequired: false,
		},
		{
			URI:          "/shops/{id}/posts",
			Method:       http.MethodGet,
			Handler:      postHandler.HandlePostsGet,
			AuthRequired: false,
		},
		{
			URI:          "/shops/{shopId}/posts/{postId}",
			Method:       http.MethodGet,
			Handler:      postHandler.HandlePostGet,
			AuthRequired: false,
		},
		{
			URI:          "/shops/{shopId}/posts/{postId}",
			Method:       http.MethodPut,
			Handler:      postHandler.HandlePostUpdate,
			AuthRequired: false,
		},
		{
			URI:          "/shops/{shopId}/posts/{postId}",
			Method:       http.MethodDelete,
			Handler:      postHandler.HandlePostDelete,
			AuthRequired: false,
		},
	}
	return postRoutes
}
