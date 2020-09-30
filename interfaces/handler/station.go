package handler

import (
	"encoding/json"
	"io/ioutil"
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
)

// StationHandler Stationに対するHandlerのインターフェイス
type StationHandler interface {
	HandleStationsSearch(w http.ResponseWriter, r *http.Request)
}

type stationHandler struct {
	stationUsecase usecase.StationUsecase
}

// NewStationHandler Stationデータに関するHandlerを生成
func NewStationHandler(su usecase.StationUsecase) StationHandler {
	return &stationHandler{
		stationUsecase: su,
	}
}

// HandleStationsSearch 該当する駅名情報を検索
func (sh stationHandler) HandleStationsSearch(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody stationSearchRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stations, err := sh.stationUsecase.SearchStations(requestBody.SearchWord)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, stations)
}

type stationSearchRequest struct {
	SearchWord string `json:"search_word"`
}
