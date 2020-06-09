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

// 全ての店舗情報のレコードを取得
func (sp *shopPersistence) FindAll() ([]models.Shop, error) {
	var err error

	shops := []models.Shop{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = sp.db.Debug().Model(&models.Shop{}).Find(&shops).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return shops, nil
	}
	return nil, err
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

// 指定した店舗情報のレコードを1件取得
func (sp *shopPersistence) FindByID(uid uint32) (models.Shop, error) {
	var err error

	shop := models.Shop{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = sp.db.Debug().Model(&models.Shop{}).Where("id = ?", uid).Take(&shop).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return shop, nil
	}
	// 指定したレコードがない場合
	if gorm.IsRecordNotFoundError(err) {
		return models.Shop{}, err
	}
	return models.Shop{}, err
}
