package controllers

import (
	"encoding/json"
	"io/ioutil"
	"merpochi_server/api/database"
	"merpochi_server/api/models"
	"merpochi_server/api/responses"
	"net/http"
)

// CreateUser ユーザー新規登録
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	err = db.Debug().Model(&models.User{}).Create(&user).Error
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}
