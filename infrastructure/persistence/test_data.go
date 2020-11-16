package persistence

import (
	"merpochi_server/domain/models"
	"time"
)

var users = []models.User{
	{
		ID:       1,
		Nickname: "miku",
		Email:    "miku@email.com",
		Password: "mikumiku",
	},
	{
		ID:       2,
		Nickname: "taka",
		Email:    "taka@email.com",
		Password: "takataka",
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
		CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
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
		Text:      "mikuのコメント001",
		ShopID:    1,
		UserID:    1,
		CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	},
	{
		Text:      "takaのコメント001",
		ShopID:    1,
		UserID:    2,
		CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	},
}
