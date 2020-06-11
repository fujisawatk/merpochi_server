package handler

import (
	"encoding/json"
	"io/ioutil"
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
)

// CommentHandler Commentに対するHandlerのインターフェイス
type CommentHandler interface {
	HandleCommentCreate(w http.ResponseWriter, r *http.Request)
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

type commentCreateRequest struct {
	Text   string `json:"text"`
	ShopID uint32 `json:"shop_id"`
}
