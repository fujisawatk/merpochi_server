package handler

import (
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
)

// UserHandler Userに対するHandlerのインターフェイス
type UserHandler interface {
	HandleUsersGet(w http.ResponseWriter, r *http.Request)
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

// HandleUsersGet ユーザー情報を全件取得
func (uh userHandler) HandleUsersGet(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userUsecase.GetUsers()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}
