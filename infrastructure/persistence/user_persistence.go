package persistence

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"

	"merpochi_server/util/channels.go"

	"github.com/jinzhu/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

// NewUserPersistence userPersistence構造体の宣言
func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{db}
}

// Save ユーザー情報保存のトランザクション
func (up *userPersistence) Save(user models.User) (models.User, error) {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = up.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

// 全てのユーザー情報のレコードを取得するトランザクション
func (up *userPersistence) FindAll() ([]models.User, error) {
	var err error

	users := []models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = up.db.Debug().Model(&models.User{}).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}
