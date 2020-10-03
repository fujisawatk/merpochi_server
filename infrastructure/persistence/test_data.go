package persistence

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
		Code:      "aaaa000",
		Name:      "焼鳥屋",
		Category:  "焼鳥",
		Opentime:  "17:00～24:00",
		Budget:    3000,
		Img:       "https://rimage.gnst.jp/rest/img/000000000000/0000.jpg",
		Latitude:  00.000000,
		Longitude: 00.000000,
		URL:       "https://r.gnavi.co.jp/000000000000/?ak=aaaaaaaa",
	},
}

var favorites = []models.Favorite{
	{
		UserID: 1,
		ShopID: 1,
	},
}

var comments = []models.Comment{
	{
		Text:   "コメントテストa00000",
		ShopID: 1,
		UserID: 1,
	},
}
