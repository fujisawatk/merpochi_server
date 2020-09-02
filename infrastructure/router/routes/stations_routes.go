package routes

import (
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/persistence"
	"merpochi_server/interfaces/handler"
	"merpochi_server/usecase"
	"net/http"
)

func iniStationsRoutes() []Route {
	// 依存関係を注入
	stationPersistence := persistence.NewStationPersistence(database.DB)
	stationUsecase := usecase.NewStationUsecase(stationPersistence)
	stationHandler := handler.NewStationHandler(stationUsecase)

	stationsRoutes := []Route{
		{
			URI:          "/stations/search",
			Method:       http.MethodPost,
			Handler:      stationHandler.HandleStationsSearch,
			AuthRequired: false,
		},
	}
	return stationsRoutes
}
