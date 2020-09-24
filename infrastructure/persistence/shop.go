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

// 店舗情報に紐づくコメント数を取得
func (sp *shopPersistence) FindCommentsCount(sid uint32) (uint32, error) {
	var err error
	var count uint32

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = sp.db.Debug().Model(&models.Comment{}).Where("shop_id = ?", sid).Count(&count).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return count, nil
	}
	return 0, err
}

// 店舗情報に紐づくいいね数を取得
func (sp *shopPersistence) FindFavoritesCount(sid uint32) (uint32, error) {
	var err error
	var count uint32

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = sp.db.Debug().Model(&models.Favorite{}).Where("shop_id = ?", sid).Count(&count).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return count, nil
	}
	return 0, err
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

func (sp *shopPersistence) SearchShop(code string) (models.Shop, error) {
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

func (sp *shopPersistence) FindCommentUser(uid uint32) (models.User, error) {
	var err error

	user := models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = sp.db.Debug().Model(&models.User{}).Where("id = ?", uid).First(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}
