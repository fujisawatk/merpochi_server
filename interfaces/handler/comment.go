package handler

import (
	"encoding/json"
	"io/ioutil"
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"merpochi_server/util/ctxval"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CommentHandler Commentに対するHandlerのインターフェイス
type CommentHandler interface {
	HandleCommentsGet(w http.ResponseWriter, r *http.Request)
	HandleCommentCreate(w http.ResponseWriter, r *http.Request)
	// HandleCommentUpdate(w http.ResponseWriter, r *http.Request)
	// HandleCommentDelete(w http.ResponseWriter, r *http.Request)
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

// HandleCommentsGet 指定の投稿に紐づくコメント情報を全件取得
func (ch *commentHandler) HandleCommentsGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["postId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	comments, err := ch.commentUsecase.GetComments(uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, comments)
}

// HandleCommentCreate 指定の投稿にコメントを登録
func (ch *commentHandler) HandleCommentCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["postId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// contextからユーザーID取得
	uid := ctxval.GetUserID(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody commentRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	comment, err := ch.commentUsecase.CreateComment(requestBody.Text, uid, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, comment)
}

// // HandleCommentUpdate 店舗情報ページに記載したコメントを編集
// func (ch commentHandler) HandleCommentUpdate(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	cid, err := strconv.ParseUint(vars["commentId"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	var requestBody commentRequest
// 	err = json.Unmarshal(body, &requestBody)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	rows, err := ch.commentUsecase.UpdateComment(uint32(cid), requestBody.Text)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	responses.JSON(w, http.StatusOK, rows)
// }

// // HandleCommentDelete 店舗情報ページに記載したコメントを削除
// func (ch commentHandler) HandleCommentDelete(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	cid, err := strconv.ParseUint(vars["commentId"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	err = ch.commentUsecase.DeleteComment(uint32(cid))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	w.Header().Set("Entity", fmt.Sprintf("%d", cid))
// 	responses.JSON(w, http.StatusNoContent, "")
// }

type commentRequest struct {
	Text string `json:"text"`
}
