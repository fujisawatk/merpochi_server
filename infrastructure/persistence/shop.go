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

func (sp *shopPersistence) FindFavoriteUser(sid, uid uint32) bool {
	var rs *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		rs = sp.db.Model(&models.Favorite{}).Where("user_id = ? AND shop_id = ?", uid, sid).Take(&models.Favorite{})
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

func (sp *shopPersistence) FindBookmarkedShops(uid uint32) (*[]models.Shop, error) {
	var err error
	shops := &[]models.Shop{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		query := sp.db.Table("users").
			Select("shops.*").
			Joins("inner join bookmarks on bookmarks.user_id = users.id").
			Joins("inner join shops on shops.id = bookmarks.shop_id").
			Where("users.id = ?", uid)
		err = query.Scan(shops).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return shops, nil
	}
	return &[]models.Shop{}, err
}

func (sp *shopPersistence) FindFavoritedShops(uid uint32) (*[]models.Shop, error) {
	var err error
	shops := &[]models.Shop{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		query := sp.db.Table("users").
			Select("shops.*").
			Joins("inner join favorites on favorites.user_id = users.id").
			Joins("inner join shops on shops.id = favorites.shop_id").
			Where("users.id = ?", uid)
		err = query.Scan(shops).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return shops, nil
	}
	return &[]models.Shop{}, err
}
