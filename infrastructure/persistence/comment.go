package persistence

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"

	"merpochi_server/util/channels"

	"github.com/jinzhu/gorm"
)

type commentPersistence struct {
	db *gorm.DB
}

// NewCommentPersistence commentPersistence構造体の宣言
func NewCommentPersistence(db *gorm.DB) repository.CommentRepository {
	return &commentPersistence{db}
}

// Save コメント保存
func (cp *commentPersistence) Save(comment models.Comment) (models.Comment, error) {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = cp.db.Debug().Model(&models.Comment{}).Create(&comment).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return comment, nil
	}
	return models.Comment{}, err
}
