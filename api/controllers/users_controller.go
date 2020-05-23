package controllers

import (
	"log"
	"net/http"
)

// CreateUser ユーザー新規登録
func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Print("hello public!")
}
