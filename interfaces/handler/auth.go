package handler

import (
	"encoding/json"
	"io/ioutil"
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
)

// AuthHandler ユーザー認証に対するHandlerのインターフェイス
type AuthHandler interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authUsecase usecase.AuthUsecase
}

// NewAuthHandler ユーザー認証に関するHandlerを生成
func NewAuthHandler(ua usecase.AuthUsecase) AuthHandler {
	return &authHandler{
		authUsecase: ua,
	}
}

// Login ログイン処理
func (ah authHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var token string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var requestBody authRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err = ah.authUsecase.LoginUser(requestBody.Email, requestBody.Password)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
