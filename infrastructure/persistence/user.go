package persistence

import (
	"errors"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"time"

	"merpochi_server/util/channels"

	"github.com/jinzhu/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

// NewUserPersistence userPersistence構造体の宣言
func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{db}
}

// Save ユーザー情報の保存
func (up *userPersistence) Save(user models.User) (models.User, error) {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = up.db.Model(&models.User{}).Create(&user).Error
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

// 指定したユーザー情報のレコードを1件取得
func (up *userPersistence) FindByID(uid uint32) (models.User, error) {
	var err error

	user := models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = up.db.Model(&models.User{}).Where("id = ?", uid).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	// 指定したレコードがない場合
	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("指定したユーザーは登録されていません")
	}
	return models.User{}, err
}

// 指定したユーザー情報のレコードを1件更新
func (up *userPersistence) Update(uid uint32, user models.User) (int64, error) {
	var rs *gorm.DB

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		rs = up.db.Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).UpdateColumns(
			map[string]interface{}{
				"nickname":   user.Nickname,
				"email":      user.Email,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		// RowsAffected→更新したレコード数を取得
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

// 指定したユーザー情報のレコードを1件削除
func (up *userPersistence) Delete(uid uint32) (int64, error) {
	var rs *gorm.DB

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		rs = up.db.Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		// RowsAffected→削除したレコード数を取得
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}
