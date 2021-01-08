package auto

import (
	"log"
	"merpochi_server/domain/models"
	"merpochi_server/infrastructure/database"
	"merpochi_server/util/console"
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
	err = db.Debug().Delete(&comments).Error
	if err != nil {
		log.Fatal(err)
	}
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

	err = db.Debug().Delete(&images).Error
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
		// var count uint32
		// err = db.Debug().Model(&models.Favorite{}).Where("shop_id = ?", favorite.ShopID).Count(&count).Error
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}

	for _, comment := range comments {
		// 店舗情報を保存
		err = db.Debug().Model(&models.Comment{}).Create(&comment).Error
		if err != nil {
			log.Fatal(err)
		}
	}
	// 指定した店舗IDのコメントを取得
	// shopComments := db.Debug().Model(&models.Comment{}).Where("shop_id = ?", 2).Find(&comments)
	// console.Pretty(shopComments)

	// テーブル結合で指定した店舗IDのコメントを取得
	type Results struct {
		ID          uint32
		Code        string
		Text        string
		UserID      uint32
		CommentUser string
	}

	var results []Results
	query := db.Debug().Table("shops").
		Select("shops.id, shops.code, comments.text, comments.user_id").
		Joins("left join comments on comments.shop_id = shops.id").
		Where("shops.code = ?", "a00000")
	query.Scan(&results)

	commentUser := models.User{}
	// ショップに紐付くコメントからユーザーIDを取得
	// ユーザーIDに紐付くユーザーを取得
	for i := 0; i < len(results); i++ {
		err = db.Debug().Model(&models.User{}).Where("id = ?", results[i].UserID).First(&commentUser).Error
		if err != nil {
			log.Fatal(err)
		}
		results[i].CommentUser = commentUser.Nickname
	}
	console.Pretty(results)

	for _, image := range images {
		// 店舗情報を保存
		err = db.Debug().Model(&models.Image{}).Create(&image).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}
