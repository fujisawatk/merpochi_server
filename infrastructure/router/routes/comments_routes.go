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
			URI:          "/comments",
			Method:       http.MethodPost,
			Handler:      commentHandler.HandleCommentCreate,
			AuthRequired: false,
		},
		{
			URI:          "/comments/{id}",
			Method:       http.MethodPut,
			Handler:      commentHandler.HandleCommentUpdate,
			AuthRequired: false,
		},
	}
	return commentRoutes
}
