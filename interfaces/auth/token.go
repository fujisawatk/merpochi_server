package auth

import (
	"merpochi_server/config"
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
