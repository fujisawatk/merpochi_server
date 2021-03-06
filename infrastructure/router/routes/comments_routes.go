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
	userPersistence := persistence.NewUserPersistence(database.DB)
	postPersistence := persistence.NewPostPersistence(database.DB)
	imagePersistence := persistence.NewImagePersistence(database.DB)
	commentUsecase := usecase.NewCommentUsecase(
		commentPersistence,
		userPersistence,
		postPersistence,
		imagePersistence,
	)
	commentHandler := handler.NewCommentHandler(commentUsecase)

	commentRoutes := []Route{
		{
			URI:          "/posts/{postId}/comments",
			Method:       http.MethodGet,
			Handler:      commentHandler.HandleCommentsGet,
			AuthRequired: false,
		},
		{
			URI:          "/posts/{postId}/comments",
			Method:       http.MethodPost,
			Handler:      commentHandler.HandleCommentCreate,
			AuthRequired: true,
		},
		{
			URI:          "/posts/{postId}/comments/{commentId}",
			Method:       http.MethodPut,
			Handler:      commentHandler.HandleCommentUpdate,
			AuthRequired: true,
		},
		{
			URI:          "/posts/{postId}/comments/{commentId}",
			Method:       http.MethodDelete,
			Handler:      commentHandler.HandleCommentDelete,
			AuthRequired: true,
		},
	}
	return commentRoutes
}
