package auto

import "merpochi_server/domain/models"

var users = []models.User{
	{
		Nickname: "miku",
		Email:    "miku@email.com",
		Password: "mikumiku",
	},
}

var shops = []models.Shop{
	{
		Code: "a00000",
	},
}

var favorites = []models.Favorite{
	{
		UserID: 1,
		ShopID: 1,
	},
}
