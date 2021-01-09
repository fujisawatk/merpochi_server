package persistence

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"time"

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

// 指定した店舗に紐付く投稿情報のレコードを全件取得
func (pp *postPersistence) FindAll(sid uint32) (*[]models.Post, error) {
	var err error
	posts := &[]models.Post{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = pp.db.Model(&models.Post{}).Where("shop_id = ?", sid).Find(posts).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return posts, nil
	}
	return &[]models.Post{}, err
}

// 投稿情報のレコードを1件取得
func (pp *postPersistence) FindByID(sid, pid uint32) (*models.Post, error) {
	var err error
	post := &models.Post{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = pp.db.Model(&models.Post{}).Where("shop_id = ? AND id = ?", sid, pid).Take(post).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return post, nil
	}
	return &models.Post{}, err
}

// 投稿情報のレコードを1件更新
func (pp *postPersistence) Update(post *models.Post) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		rs = pp.db.Model(&models.Post{}).Where("id = ?", post.ID).Take(&models.Post{}).UpdateColumns(
			map[string]interface{}{
				"text":       post.Text,
				"rating":     post.Rating,
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

// 投稿情報のレコードを1件削除
func (pp *postPersistence) Delete(pid uint32) error {
	var rs *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		rs = pp.db.Model(&models.Post{}).Where("id = ?", pid).Take(&models.Post{}).Delete(&models.Post{})
		ch <- true
	}(done)
	if channels.OK(done) {
		return nil
	}
	return rs.Error
}
