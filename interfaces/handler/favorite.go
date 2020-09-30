package handler

import (
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"merpochi_server/util/ctxval"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// FavoriteHandler Userに対するHandlerのインターフェイス
type FavoriteHandler interface {
	HandleFavoritesGet(w http.ResponseWriter, r *http.Request)
	HandleFavoriteCreate(w http.ResponseWriter, r *http.Request)
	HandleFavoriteDelete(w http.ResponseWriter, r *http.Request)
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

// HandleFavoriteGet 指定の店舗に紐付くお気に入り情報を取得
func (fh favoriteHandler) HandleFavoritesGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["shopId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	favorites, err := fh.favoriteUsecase.GetFavorites(uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, favorites)
}

// HandleFavoriteCreate お気に入りを登録
func (fh favoriteHandler) HandleFavoriteCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["shopId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid := ctxval.GetUserID(r)

	favorite, err := fh.favoriteUsecase.CreateFavorite(uint32(sid), uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, favorite)
}

// HandleFavoriteDelete お気に入りを解除
func (fh favoriteHandler) HandleFavoriteDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["shopId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid := ctxval.GetUserID(r)

	err = fh.favoriteUsecase.DeleteFavorite(uint32(sid), uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
