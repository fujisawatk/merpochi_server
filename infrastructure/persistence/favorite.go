package persistence

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"

	"merpochi_server/util/channels"

	"github.com/jinzhu/gorm"
)

type favoritePersistence struct {
	db *gorm.DB
}

// NewFavoritePersistence favoritePersistence構造体の宣言
func NewFavoritePersistence(db *gorm.DB) repository.FavoriteRepository {
	return &favoritePersistence{db}
}

// Save お気に入り登録
func (fp *favoritePersistence) Save(favorite models.Favorite) error {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = fp.db.Debug().Model(&models.Favorite{}).Create(&favorite).Error
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

// Search 指定した店舗IDが登録されているレコード数を取得
func (fp *favoritePersistence) Search(sid uint32) (uint32, error) {
	var err error
	var count uint32
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		// SELECT count(*) FROM favorites WHERE shop_id = sid; (count)
		err = fp.db.Debug().Model(&models.Favorite{}).Where("shop_id = ?", sid).Count(&count).Error
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
