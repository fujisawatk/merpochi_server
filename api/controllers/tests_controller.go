package controllers

import (
	"log"
	"net/http"
)

// Public テスト用
func Public(w http.ResponseWriter, r *http.Request) {
	log.Print("hello public!")
}
