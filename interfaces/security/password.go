package security

import "golang.org/x/crypto/bcrypt"

// Hash パスワードのハッシュ化
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
