package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PostHandler Userに対するHandlerのインターフェイス
type PostHandler interface {
	HandlePostCreate(w http.ResponseWriter, r *http.Request)
	HandlePostsGet(w http.ResponseWriter, r *http.Request)
	HandlePostGet(w http.ResponseWriter, r *http.Request)
	HandlePostUpdate(w http.ResponseWriter, r *http.Request)
	HandlePostDelete(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	postUsecase usecase.PostUsecase
}

// NewPostHandler Userデータに関するHandlerを生成
func NewPostHandler(pu usecase.PostUsecase) PostHandler {
	return &postHandler{
		postUsecase: pu,
	}
}

// HandlePostCreate 投稿情報を登録
func (ph *postHandler) HandlePostCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody postCreateRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post, err := ph.postUsecase.CreatePost(requestBody.Text, requestBody.Rating, requestBody.UserID, uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, post)
}

// HandlePostsGet 指定した店舗に紐付く投稿情報を全件取得
func (ph *postHandler) HandlePostsGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	posts, err := ph.postUsecase.GetPosts(uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

// HandlePostGet 投稿情報を1件取得
func (ph *postHandler) HandlePostGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["shopId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	pid, err := strconv.ParseUint(vars["postId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	post, err := ph.postUsecase.GetPost(uint32(sid), uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, post)
}

// HandlePostUpdate 投稿情報を1件更新
func (ph *postHandler) HandlePostUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["postId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody postUpdateRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	rows, err := ph.postUsecase.UpdatePost(uint32(pid), requestBody.Rating, requestBody.Text)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, rows)
}

// HandlePostDelete 投稿情報を1件削除
func (ph *postHandler) HandlePostDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["postId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = ph.postUsecase.DeletePost(uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}

type postCreateRequest struct {
	Text   string `json:"text"`
	Rating uint32 `json:"rating"`
	UserID uint32 `json:"user_id"`
}

type postUpdateRequest struct {
	Text   string `json:"text"`
	Rating uint32 `json:"rating"`
}
