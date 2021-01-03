package healthcheck

import (
	"io"
	"net/http"
)

// HandleHealthCheck ヘルスチェック用
func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}
