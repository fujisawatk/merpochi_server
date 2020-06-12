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

// 指定した店舗のコメント情報を取得（店舗情報はフロント側の外部APIから取得し表示）
func (sp *shopPersistence) FindByID(sid uint32) ([]models.Comment, error) {
	var results []models.Comment

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		query := sp.db.Debug().Table("shops").
			Select("comments.id, comments.text").
			Joins("inner join comments on comments.shop_id = shops.id").
			Where("shops.id = ?", sid)
		query.Scan(&results)
		if len(results) == 0 {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return results, nil
	}
	return []models.Comment{}, errors.New("no comment")
}
