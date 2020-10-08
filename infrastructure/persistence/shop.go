package persistence

import (
	"errors"
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

// 店舗情報に紐づくコメント数を取得
func (sp *shopPersistence) FindCommentsCount(sid uint32) uint32 {
	var count uint32

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		sp.db.Debug().Model(&models.Comment{}).Where("shop_id = ?", sid).Count(&count)
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
		sp.db.Debug().Model(&models.Favorite{}).Where("shop_id = ?", sid).Count(&count)
		ch <- true
	}(done)
	if channels.OK(done) {
		return count
	}
	return 0
}

// 店舗情報を保存
func (sp *shopPersistence) Save(shop models.Shop) (models.Shop, error) {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = sp.db.Debug().Model(&models.Shop{}).Create(&shop).Error
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

func (sp *shopPersistence) Search(code string) (models.Shop, error) {
	var err error

	shop := models.Shop{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = sp.db.Debug().Model(&models.Shop{}).Where("code = ?", code).Take(&shop).Error
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

func (sp *shopPersistence) FindCommentedShops(uid uint32) ([]models.Shop, error) {
	var shops []models.Shop
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		query := sp.db.Debug().Table("users").
			Select("shops.*").
			Joins("inner join comments on comments.user_id = users.id").
			Joins("inner join shops on shops.id = comments.shop_id").
			Where("users.id = ?", uid)
		err = query.Scan(&shops).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return shops, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return []models.Shop{}, errors.New("shop not found")
	}
	return []models.Shop{}, err
}

func (sp *shopPersistence) FindFavoritedShops(uid uint32) ([]models.Shop, error) {
	var shops []models.Shop
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		query := sp.db.Debug().Table("users").
			Select("shops.*").
			Joins("inner join favorites on favorites.user_id = users.id").
			Joins("inner join shops on shops.id = favorites.shop_id").
			Where("users.id = ?", uid)
		err = query.Scan(&shops).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return shops, nil
	}
	return []models.Shop{}, err
}
