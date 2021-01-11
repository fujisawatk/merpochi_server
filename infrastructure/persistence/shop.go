package persistence

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"

	"merpochi_server/util/channels"

	"github.com/jinzhu/gorm"
)

type shopPersistence struct {
	db *gorm.DB
}

// NewShopPersistence shopPersistence構造体の宣言
func NewShopPersistence(db *gorm.DB) repository.ShopRepository {
	return &shopPersistence{db}
}

// 評価が4以上である投稿数を取得
func (sp *shopPersistence) FindPostsCount(pid uint32) uint32 {
	var count uint32

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		sp.db.Model(&models.Post{}).Where("shop_id = ? AND rating >= ?", pid, 4).Count(&count)
		ch <- true
	}(done)
	if channels.OK(done) {
		return count
	}
	return 0
}

// 店舗情報に紐づくお気に入り数を取得
func (sp *shopPersistence) FindFavoritesCount(sid uint32) uint32 {
	var count uint32

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		sp.db.Model(&models.Favorite{}).Where("shop_id = ?", sid).Count(&count)
		ch <- true
	}(done)
	if channels.OK(done) {
		return count
	}
	return 0
}

// 指定した店舗のブックマーク数を取得
func (sp *shopPersistence) FindBookmarksCount(sid uint32) uint32 {
	var count uint32

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		sp.db.Model(&models.Bookmark{}).Where("shop_id = ?", sid).Count(&count)
		ch <- true
	}(done)
	if channels.OK(done) {
		return count
	}
	return 0
}

func (sp *shopPersistence) FindBookmarkUser(sid, uid uint32) bool {
	var rs *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		rs = sp.db.Model(&models.Bookmark{}).Where("user_id = ? AND shop_id = ?", uid, sid).Take(&models.Bookmark{})
		if rs.Error != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return true
	}
	return false
}

// 店舗情報を保存
func (sp *shopPersistence) Save(shop models.Shop) (models.Shop, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = sp.db.Model(&models.Shop{}).Create(&shop).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return shop, nil
	}
	return models.Shop{}, err
}

func (sp *shopPersistence) FindByCode(code string) (*models.Shop, error) {
	var err error
	shop := &models.Shop{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = sp.db.Model(&models.Shop{}).Where("code = ?", code).Take(shop).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return shop, nil
	}
	return &models.Shop{}, err
}
