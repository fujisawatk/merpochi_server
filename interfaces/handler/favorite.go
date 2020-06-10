package handler

import (
	"encoding/json"
	"io/ioutil"
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// FavoriteHandler Userに対するHandlerのインターフェイス
type FavoriteHandler interface {
	HandleFavoriteCreate(w http.ResponseWriter, r *http.Request)
	// HandleFavoriteDelete(w http.ResponseWriter, r *http.Request)
}

type favoriteHandler struct {
	favoriteUsecase usecase.FavoriteUsecase
}

// NewFavoriteHandler Userデータに関するHandlerを生成
func NewFavoriteHandler(fu usecase.FavoriteUsecase) FavoriteHandler {
	return &favoriteHandler{
		favoriteUsecase: fu,
	}
}

// HandleFavoriteCreate お気に入りを登録
func (fh favoriteHandler) HandleFavoriteCreate(w http.ResponseWriter, r *http.Request) {
	// お気に入りした店舗情報ページのID取得
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// リクエストボディからユーザーID取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody favoriteRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// 成功したら指定ページのいいね総数を返す
	count, err := fh.favoriteUsecase.CreateFavorite(uint32(sid), requestBody.UserID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, count)
}

type favoriteRequest struct {
	UserID uint32 `json:"user_id"`
}
