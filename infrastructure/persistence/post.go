package persistence

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"

	"merpochi_server/util/channels"

	"github.com/jinzhu/gorm"
)

type postPersistence struct {
	db *gorm.DB
}

// NewPostPersistence postPersistence構造体の宣言
func NewPostPersistence(db *gorm.DB) repository.PostRepository {
	return &postPersistence{db}
}

// Save 投稿情報の保存
func (pp *postPersistence) Save(post *models.Post) error {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = pp.db.Model(&models.Post{}).Create(&post).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return nil
	}
	return err
}
