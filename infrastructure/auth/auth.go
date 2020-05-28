package auth

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/security"
	"merpochi_server/infrastructure/database"

	"github.com/jinzhu/gorm"
	"merpochi_server/util/channels.go"
)

// SignIn 既存ユーザーか否か確認
func SignIn(email, password string) (string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
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
