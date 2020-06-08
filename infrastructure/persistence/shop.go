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
