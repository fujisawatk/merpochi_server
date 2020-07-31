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
		ID:        1,
		Code:      "gcrk447",
		Name:      "焼鳥屋 鳥貴族 鶴見西口店",
		Category:  "焼鳥",
		Opentime:  "17:00～24:00",
		Budget:    2000,
		Img:       "https://rimage.gnst.jp/rest/img/b1639pht0000/t_0omo.jpg",
		Latitude:  35.509054,
		Longitude: 139.674982,
		URL:       "https://r.gnavi.co.jp/b1639pht0000/?ak=uGZLtrurN0Ig0vhqZz1VLGDDPH0WafuFnTh%2Bu4ABpI8%3D",
	},
	{
		ID:        2,
		Code:      "gdyd408",
		Name:      "焼鳥屋 鳥貴族 鶴見東口店",
		Category:  "焼鳥",
		Opentime:  "17:00～24:00",
		Budget:    2000,
		Img:       "https://rimage.gnst.jp/rest/img/nnb7n5up0000/t_0oly.jpg",
		Latitude:  35.508283,
		Longitude: 139.677655,
		URL:       "https://r.gnavi.co.jp/nnb7n5up0000/?ak=uGZLtrurN0Ig0vhqZz1VLGDDPH0WafuFnTh%2Bu4ABpI8%3D",
	},
}

var favorites = []models.Favorite{
	{
		ID:     1,
		UserID: 1,
		ShopID: 1,
	},
	{
		ID:     2,
		UserID: 1,
		ShopID: 2,
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
	{
		ID:     4,
		Text:   "コメントテストa00000-3",
		ShopID: 2,
		UserID: 1,
	},
}
