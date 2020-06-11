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

// CommentHandler Commentに対するHandlerのインターフェイス
type CommentHandler interface {
	HandleCommentCreate(w http.ResponseWriter, r *http.Request)
	HandleCommentUpdate(w http.ResponseWriter, r *http.Request)
}

type commentHandler struct {
	commentUsecase usecase.CommentUsecase
}

// NewCommentHandler Commentデータに関するHandlerを生成
func NewCommentHandler(cu usecase.CommentUsecase) CommentHandler {
	return &commentHandler{
		commentUsecase: cu,
	}
}

// HandleCommentCreate 店舗情報ページにコメントを登録
func (ch commentHandler) HandleCommentCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody commentCreateRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	comment, err := ch.commentUsecase.CreateComment(requestBody.Text, requestBody.ShopID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, comment)
}

// HandleCommentUpdate 店舗情報ページに記載したコメントを編集
func (ch commentHandler) HandleCommentUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody commentUpdateRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	rows, err := ch.commentUsecase.UpdateComment(uint32(cid), requestBody.Text)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, rows)
}

type commentCreateRequest struct {
	Text   string `json:"text"`
	ShopID uint32 `json:"shop_id"`
}

type commentUpdateRequest struct {
	Text string `json:"text"`
}
