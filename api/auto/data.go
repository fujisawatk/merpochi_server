package auto

import "merpochi_server/domain/models"

var users = []models.User{
	{
		ID:       1,
		Nickname: "miku",
		Email:    "miku@email.com",
		Password: "mikumiku",
	},
	{
		ID:       2,
		Nickname: "fuji",
		Email:    "fuji@email.com",
		Password: "fujifuji",
	},
}

var shops = []models.Shop{
	{
		ID:       1,
		Code:     "gcrk447",
		Name:     "焼鳥屋 鳥貴族 鶴見西口店",
		Category: "焼鳥",
	},
	{
		ID:       2,
		Code:     "gdyd408",
		Name:     "焼鳥屋 鳥貴族 鶴見東口店",
		Category: "焼鳥",
	},
}

var favorites = []models.Favorite{
	{
		ID:     1,
		UserID: 1,
		ShopID: 1,
	},
}

var comments = []models.Comment{
	{
		ID:     1,
		Text:   "コメントテストa00000",
		ShopID: 1,
		UserID: 1,
	},
	{
		ID:     2,
		Text:   "コメントテストb11111",
		ShopID: 2,
		UserID: 2,
	},
	{
		ID:     3,
		Text:   "コメントテストa00000-2",
		ShopID: 2,
		UserID: 1,
	},
}
