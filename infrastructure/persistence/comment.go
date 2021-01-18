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

func (cp *commentPersistence) FindAll(pid uint32) (*[]models.Comment, error) {
	var err error
	comments := &[]models.Comment{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err := cp.db.Model(&models.Comment{}).Where("post_id = ?", pid).Find(comments).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return comments, nil
	}
	return &[]models.Comment{}, err
}

// Save コメント保存
func (cp *commentPersistence) Save(comment *models.Comment) error {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = cp.db.Model(&models.Comment{}).Create(comment).Error
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

// 指定したコメントのレコードを1件更新
func (cp *commentPersistence) Update(cid uint32, comment *models.Comment) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		rs = cp.db.Model(&models.Comment{}).Where("id = ?", cid).Take(&models.Comment{}).UpdateColumns(
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

// 指定したコメントのレコードを1件削除
func (cp *commentPersistence) Delete(cid uint32) error {
	var rs *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		rs = cp.db.Model(&models.Comment{}).Where("id = ?", cid).Take(&models.Comment{}).Delete(&models.Comment{})
		ch <- true
	}(done)
	if channels.OK(done) {
		return nil
	}
	return rs.Error
}

func (cp *commentPersistence) FindByUserID(uid uint32) (*models.User, error) {
	var err error
	user := &models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = cp.db.Model(&models.User{}).Where("id = ?", uid).First(user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return &models.User{}, err
}

// 投稿情報に紐づくコメント数を取得
func (cp *commentPersistence) CountByPostID(pid uint32) uint32 {
	var count uint32
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		cp.db.Model(&models.Comment{}).Where("post_id = ?", pid).Count(&count)
		ch <- true
	}(done)
	if channels.OK(done) {
		return count
	}
	return 0
}
