package auth

import (
	"context"
	"fmt"
	"merpochi_server/config"
	"merpochi_server/util/ctxval"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken トークン作成
func CreateToken(userID uint32) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SECRETKEY)
}

// TokenValid トークンのバリデーションチェック
func TokenValid(r *http.Request) (context.Context, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.SECRETKEY, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx := ctxval.SetUserID(r, claims["user_id"].(float64))
		return ctx, nil
	}
	fmt.Println(err)
	return nil, err
}

// ExtractToken リクエスト情報からトークン取得
func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	// Authorization形式例→"Basic YWxhZGRpbjpvcGVuc2VzYW1l(認証トークン)"
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
