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

// UserHandler Userに対するHandlerのインターフェイス
type UserHandler interface {
	HandleUserCreate(w http.ResponseWriter, r *http.Request)
	HandleUserGet(w http.ResponseWriter, r *http.Request)
	HandleUserUpdate(w http.ResponseWriter, r *http.Request)
	HandleUserDelete(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler Userデータに関するHandlerを生成
func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

// HandleUserCreate ユーザー情報を登録
func (uh userHandler) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody userCreateRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user, err := uh.userUsecase.CreateUser(requestBody.Nickname, requestBody.Email, requestBody.Password, requestBody.Genre)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}

// HandleUserGet ユーザー情報を1件取得
func (uh userHandler) HandleUserGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err := uh.userUsecase.GetUser(uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

// HandleUserUpdate ユーザー情報を1件更新
func (uh userHandler) HandleUserUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody userUpdateRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	rows, err := uh.userUsecase.UpdateUser(uint32(uid), requestBody.Nickname, requestBody.Email)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, rows)
}

// HandleUserDelete ユーザー情報を1件削除
func (uh userHandler) HandleUserDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = uh.userUsecase.DeleteUser(uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

type userCreateRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Genre    string `json:"genre"`
}

type userUpdateRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}
