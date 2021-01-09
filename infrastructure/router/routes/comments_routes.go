package routes

import (
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/persistence"
	"merpochi_server/interfaces/handler"
	"merpochi_server/usecase"
	"net/http"
)

func iniCommentsRoutes() []Route {
	// 依存関係を注入
	commentPersistence := persistence.NewCommentPersistence(database.DB)
	commentUsecase := usecase.NewCommentUsecase(commentPersistence)
	commentHandler := handler.NewCommentHandler(commentUsecase)

	commentRoutes := []Route{
		{
			URI:          "/posts/{postId}/comments",
			Method:       http.MethodGet,
			Handler:      commentHandler.HandleCommentsGet,
			AuthRequired: false,
		},
		// {
		// 	URI:          "/shops/{shopId}/comments",
		// 	Method:       http.MethodPost,
		// 	Handler:      commentHandler.HandleCommentCreate,
		// 	AuthRequired: true,
		// },
		// {
		// 	URI:          "/shops/{shopId}/comments/{commentId}",
		// 	Method:       http.MethodPut,
		// 	Handler:      commentHandler.HandleCommentUpdate,
		// 	AuthRequired: true,
		// },
		// {
		// 	URI:          "/shops/{shopId}/comments/{commentId}",
		// 	Method:       http.MethodDelete,
		// 	Handler:      commentHandler.HandleCommentDelete,
		// 	AuthRequired: true,
		// },
	}
	return commentRoutes
}
