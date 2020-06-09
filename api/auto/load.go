package auto

import (
	"log"
	"merpochi_server/domain/models"
	"merpochi_server/infrastructure/database"
	"merpochi_server/util/security"
)

// TestLoad テストデータの読込
func TestLoad() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().Delete(&users).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Delete(&shops).Error
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		var hashedPassword []byte
		hashedPassword, err = security.Hash(user.Password)
		if err != nil {
			log.Fatal(err)
		}
		user.Password = string(hashedPassword)

		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, shop := range shops {
		err = db.Debug().Model(&models.Shop{}).Create(&shop).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}
