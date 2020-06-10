package auto

import (
	"log"
	"merpochi_server/domain/models"
	"merpochi_server/infrastructure/database"
	"merpochi_server/util/security"
)

// TestLoad テストデータの読込
func TestLoad() {
	// DB接続
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブルリセット
	err = db.Debug().Delete(&users).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Delete(&shops).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Delete(&favorites).Error
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		// パスワードのハッシュ化
		var hashedPassword []byte
		hashedPassword, err = security.Hash(user.Password)
		if err != nil {
			log.Fatal(err)
		}
		user.Password = string(hashedPassword)

		// ユーザー情報の保存
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, shop := range shops {
		// 店舗情報を保存
		err = db.Debug().Model(&models.Shop{}).Create(&shop).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, favorite := range favorites {
		// 指定したページのお気に入り登録
		err = db.Debug().Model(&models.Favorite{}).Create(&favorite).Error
		if err != nil {
			log.Fatal(err)
		}
		// 指定した店舗IDのレコード数を取得
		var count uint32
		err = db.Debug().Model(&models.Favorite{}).Where("shop_id = ?", favorite.ShopID).Count(&count).Error
		if err != nil {
			log.Fatal(err)
		}
		log.Println(count)
	}
}
