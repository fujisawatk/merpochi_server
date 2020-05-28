package persistence

import (
	"errors"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"time"

	"merpochi_server/util/channels.go"

	"github.com/jinzhu/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

// NewUserPersistence userPersistence構造体の宣言
func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{db}
}

// Save ユーザー情報保存のトランザクション
func (up *userPersistence) Save(user models.User) (models.User, error) {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		// 登録メールアドレスの重複検証
		err = up.db.Debug().Model(models.User{}).Where("email = ?", user.Email).Take(&user).Error
		if err == nil {
			ch <- false
			return
		}

		err = up.db.Debug().Model(&models.User{}).Create(&user).Error
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

// 全てのユーザー情報のレコードを取得するトランザクション
func (up *userPersistence) FindAll() ([]models.User, error) {
	var err error

	users := []models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = up.db.Debug().Model(&models.User{}).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

// 指定したユーザー情報のレコードを1件取得
func (up *userPersistence) FindByID(uid uint32) (models.User, error) {
	var err error

	user := models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = up.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&user).Error
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
		rs = up.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).UpdateColumns(
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
		rs = up.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})
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

// SearchUser 指定メールアドレスのユーザーを検索(無い場合にtrue)
func (up *userPersistence) SearchUser(email string) error {
	var err error

	user := models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = up.db.Debug().Model(&models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- true
			return
		}
		ch <- false
	}(done)
	if channels.OK(done) {
		return nil
	}
	return errors.New("このメールアドレスは既に使用されています")
}
