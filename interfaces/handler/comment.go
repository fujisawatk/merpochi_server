package handler

import (
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CommentHandler Commentに対するHandlerのインターフェイス
type CommentHandler interface {
	HandleCommentsGet(w http.ResponseWriter, r *http.Request)
	// HandleCommentCreate(w http.ResponseWriter, r *http.Request)
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

// HandleCommentsGet 指定の店舗に紐づくコメント情報を全て取得
func (ch commentHandler) HandleCommentsGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["shopId"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	comments, err := ch.commentUsecase.GetComments(uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, comments)
}

// HandleCommentCreate 店舗情報ページにコメントを登録
// func (ch commentHandler) HandleCommentCreate(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	sid, err := strconv.ParseUint(vars["shopId"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	uid := ctxval.GetUserID(r)

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

// 	comment, err := ch.commentUsecase.CreateComment(requestBody.Text, uint32(sid), uid)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	responses.JSON(w, http.StatusCreated, comment)
// }

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

// type commentRequest struct {
// 	Text string `json:"text"`
// }
