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
func (bp *bookmarkPersistence) Save(bookmark *models.Bookmark) (*models.Bookmark, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		// 重複チェック
		result := bp.db.Model(&models.Bookmark{}).Where("user_id = ? AND shop_id = ?", bookmark.UserID, bookmark.ShopID).Take(&models.Bookmark{})
		if result.RowsAffected > 0 {
			err = errors.New("bookmark registered")
			ch <- false
			return
		}
		err = bp.db.Model(&models.Bookmark{}).Create(&bookmark).Error
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

// Delete ブックマーク解除
func (bp *bookmarkPersistence) Delete(sid uint32, uid uint32) error {
	var rs *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		// 存在チェック
		rs = bp.db.Model(&models.Bookmark{}).Where("user_id = ? AND shop_id = ?", uid, sid).Take(&models.Bookmark{})
		if rs.Error != nil {
			ch <- false
			return
		}
		// 削除処理
		rs = bp.db.Model(&models.Bookmark{}).Where("user_id = ? and shop_id = ?", uid, sid).Delete(&models.Bookmark{})
		ch <- true
	}(done)
	if channels.OK(done) {
		return nil
	}
	return rs.Error
}

// 指定した店舗のブックマーク数を取得
func (bp *bookmarkPersistence) CountByShopID(sid uint32) uint32 {
	var count uint32

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		bp.db.Model(&models.Bookmark{}).Where("shop_id = ?", sid).Count(&count)
		ch <- true
	}(done)
	if channels.OK(done) {
		return count
	}
	return 0
}
