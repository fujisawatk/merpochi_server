package routes

import (
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/persistence"
	"merpochi_server/interfaces/handler"
	"merpochi_server/usecase"
	"net/http"
)

func iniImagesRoutes() []Route {
	imagePersistence := persistence.NewImagePersistence(database.DB)
	imageUsecase := usecase.NewImageUsecase(imagePersistence)
	imageHandler := handler.NewImageHandler(imageUsecase)

	imagesRoutes := []Route{
		{
			URI:          "/users/{id}/image",
			Method:       http.MethodPost,
			Handler:      imageHandler.HandleImageCreate,
			AuthRequired: false,
		},
		{
			URI:          "/users/{id}/image",
			Method:       http.MethodGet,
			Handler:      imageHandler.HandleImageGet,
			AuthRequired: false,
		},
		{
			URI:          "/users/{id}/image",
			Method:       http.MethodPut,
			Handler:      imageHandler.HandleImageUpdate,
			AuthRequired: false,
		},
	}
	return imagesRoutes
}
