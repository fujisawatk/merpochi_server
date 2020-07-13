package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
	"strings"
)

// AuthHandler ユーザー認証に対するHandlerのインターフェイス
type AuthHandler interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleVerify(w http.ResponseWriter, r *http.Request)
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

	user, err := ah.authUsecase.LoginUser(requestBody.Email, requestBody.Password)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

// HandleVerify ユーザー確認（リクエストが来たら、ユーザーデータを返す）
func (ah authHandler) HandleVerify(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) == 2 {
		authToken := bearerToken[1]
		user, err := ah.authUsecase.VerifyUser(authToken)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, user)
	} else {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("トークンの形式が不正です"))
		return
	}
}

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
