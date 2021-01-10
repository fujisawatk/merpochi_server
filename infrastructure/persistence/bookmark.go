package persistence

import (
	"errors"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"

	"merpochi_server/util/channels"

	"github.com/jinzhu/gorm"
)

type bookmarkPersistence struct {
	db *gorm.DB
}

// NewBookmarkPersistence bookmarkPersistence構造体の宣言
func NewBookmarkPersistence(db *gorm.DB) repository.BookmarkRepository {
	return &bookmarkPersistence{db}
}

// Save お気に入り登録
func (fp *bookmarkPersistence) Save(bookmark *models.Bookmark) (*models.Bookmark, error) {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		// 重複チェック
		result := fp.db.Model(&models.Bookmark{}).Where("user_id = ? AND shop_id = ?", bookmark.UserID, bookmark.ShopID).Take(&models.Bookmark{})
		if result.RowsAffected > 0 {
			err = errors.New("bookmark registered")
			ch <- false
			return
		}
		err = fp.db.Model(&models.Bookmark{}).Create(&bookmark).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return bookmark, nil
	}
	return &models.Bookmark{}, err
}
