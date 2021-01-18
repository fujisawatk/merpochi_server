package persistence

import (
	"errors"
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

// 指定した店舗のいいね情報を取得
func (fp *favoritePersistence) FindAll(sid uint32) ([]models.Favorite, error) {
	var err error
	var favorites []models.Favorite

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		query := fp.db.Table("shops").
			Select("favorites.*").
			Joins("inner join favorites on favorites.shop_id = shops.id").
			Where("shops.id = ?", sid)
		err = query.Scan(&favorites).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return favorites, nil
	}
	return []models.Favorite{}, err
}

// Save お気に入り登録
func (fp *favoritePersistence) Save(favorite models.Favorite) (models.Favorite, error) {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		// 重複チェック
		result := fp.db.Model(&models.Favorite{}).Where("user_id = ? AND shop_id = ?", favorite.UserID, favorite.ShopID).Take(&models.Favorite{})
		if result.RowsAffected > 0 {
			err = errors.New("favorite registered")
			ch <- false
			return
		}
		err = fp.db.Model(&models.Favorite{}).Create(&favorite).Error
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

// Delete お気に入り解除
func (fp *favoritePersistence) Delete(sid uint32, uid uint32) (int64, error) {
	var rs *gorm.DB

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		// 存在チェック
		rs = fp.db.Model(&models.Favorite{}).Where("user_id = ? AND shop_id = ?", uid, sid).Take(&models.Favorite{})
		if rs.Error != nil {
			ch <- false
			return
		}
		// 削除処理
		rs = fp.db.Model(&models.Favorite{}).Where("user_id = ? and shop_id = ?", uid, sid).Delete(&models.Favorite{})
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, errors.New("favorite not found")
}

func (fp *favoritePersistence) FindFavoriteUser(uid uint32) (models.User, error) {
	var err error

	user := models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = fp.db.Model(&models.User{}).Where("id = ?", uid).Take(&user).Error
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

// 店舗情報に紐づくお気に入り数を取得
func (fp *favoritePersistence) FindFavoritesCount(sid uint32) uint32 {
	var count uint32

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		fp.db.Model(&models.Favorite{}).Where("shop_id = ?", sid).Count(&count)
		ch <- true
	}(done)
	if channels.OK(done) {
		return count
	}
	return 0
}
