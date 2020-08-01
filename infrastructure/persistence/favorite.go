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
func (fp *favoritePersistence) Save(favorite models.Favorite) (models.Favorite, error) {
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
		return favorite, nil
	}
	return models.Favorite{}, err
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

// Delete お気に入り解除
func (fp *favoritePersistence) Delete(sid uint32, uid uint32) error {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = fp.db.Debug().Model(&models.Favorite{}).Where("user_id = ? and shop_id = ?", uid, sid).Delete(&models.Favorite{}).Error
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

func (fp *favoritePersistence) FindFavoriteUser(uid uint32) (models.User, error) {
	var err error

	user := models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = fp.db.Debug().Model(&models.User{}).Where("id = ?", uid).First(&user).Error
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
