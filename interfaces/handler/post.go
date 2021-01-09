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

// PostHandler Userに対するHandlerのインターフェイス
type PostHandler interface {
	HandlePostCreate(w http.ResponseWriter, r *http.Request)
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

type postCreateRequest struct {
	Text   string `json:"text"`
	Rating uint32 `json:"rating"`
	UserID uint32 `json:"user_id"`
}
