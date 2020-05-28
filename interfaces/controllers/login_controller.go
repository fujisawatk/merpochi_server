package controllers

import (
	"encoding/json"
	"io/ioutil"
	"merpochi_server/domain/models"
	"merpochi_server/interfaces/auth"
	"merpochi_server/interfaces/responses"
	"merpochi_server/interfaces/validations"
	"net/http"
)

// Login ログイン処理
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	errs := validations.UserLoginValidate(&user)
	if len(errs) > 1 {
		responses.ERRORS(w, http.StatusBadRequest, errs)
		return
	}

	token, err := auth.SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}
