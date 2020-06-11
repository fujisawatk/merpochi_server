package persistence

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"time"

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

// 指定したコメントのレコードを1件更新
func (cp *commentPersistence) Update(cid uint32, comment models.Comment) (int64, error) {
	var rs *gorm.DB

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		rs = cp.db.Debug().Model(&models.Comment{}).Where("id = ?", cid).Take(&models.Comment{}).UpdateColumns(
			map[string]interface{}{
				"text":       comment.Text,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		// RowsAffected→更新したレコード数を取得
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}
