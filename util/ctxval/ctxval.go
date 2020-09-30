package ctxval

import (
	"context"
	"net/http"
)

type key int

const authKey key = iota

// GetUserID ログインユーザーIDをcontextから取得
func GetUserID(r *http.Request) uint32 {
	uid := r.Context().Value(authKey)
	return uint32(uid.(float64))
}

// SetUserID ログインユーザーIDをcontextに設定
func SetUserID(r *http.Request, uid float64) context.Context {
	return context.WithValue(r.Context(), authKey, uid)
}
