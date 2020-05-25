package auto

import (
	"log"
	"merpochi_server/domain/models"
	"merpochi_server/infrastructure/database"
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

	for _, users := range users {
		err = db.Debug().Model(&models.User{}).Create(&users).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}
