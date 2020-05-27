package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON レスポンスボディの書込
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR エラー時のレスポンス
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

// ERRORS バリデーションエラー時のレスポンス
func ERRORS(w http.ResponseWriter, statusCode int, errs []string) {
	if errs != nil {
		JSON(w, statusCode, errs)
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
