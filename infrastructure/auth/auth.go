package auth

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/util/security"

	"merpochi_server/util/channels"

	"github.com/jinzhu/gorm"
)

type authPersistence struct {
	db *gorm.DB
}

// NewAuthPersistence authPersistence構造体の宣言
func NewAuthPersistence(db *gorm.DB) repository.AuthRepository {
	return &authPersistence{db}
}

// SignIn 既存ユーザーか否か確認
func (ap *authPersistence) SignIn(email, password string) (string, error) {
	user := models.User{}
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = ap.db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}

		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return CreateToken(user.ID)
	}
	return "", err
}
