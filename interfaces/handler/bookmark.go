package handler

import (
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"merpochi_server/util/ctxval"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// BookmarkHandler Bookmarkに対するHandlerのインターフェイス
type BookmarkHandler interface {
	HandleBookmarkCreate(w http.ResponseWriter, r *http.Request)
	HandleBookmarkDelete(w http.ResponseWriter, r *http.Request)
}

type bookmarkHandler struct {
	bookmarkUsecase usecase.BookmarkUsecase
}

// NewBookmarkHandler Bookmarkデータに関するHandlerを生成
func NewBookmarkHandler(bu usecase.BookmarkUsecase) BookmarkHandler {
	return &bookmarkHandler{
		bookmarkUsecase: bu,
	}
}

// HandleBookmarkCreate お気に入りを登録
func (bh *bookmarkHandler) HandleBookmarkCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["shopId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid := ctxval.GetUserID(r)

	bookmark, err := bh.bookmarkUsecase.CreateBookmark(uint32(sid), uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, bookmark)
}

// HandleBookmarkDelete ブックマークを解除
func (bh *bookmarkHandler) HandleBookmarkDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["shopId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid := ctxval.GetUserID(r)

	err = bh.bookmarkUsecase.DeleteBookmark(uint32(sid), uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
