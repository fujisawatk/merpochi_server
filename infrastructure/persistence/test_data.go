package persistence

import (
	"merpochi_server/domain/models"
	"time"
)

var users = []models.User{
	{
		Nickname: "miku",
		Email:    "miku@email.com",
		Password: "mikumiku",
	},
	{
		Nickname: "taka",
		Email:    "taka@email.com",
		Password: "takataka",
	},
	{
		Nickname: "enako",
		Email:    "enako@email.com",
		Password: "enaena",
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
	{
		Code:      "bbbb000",
		Name:      "海鮮居酒屋",
		Category:  "海鮮",
		Opentime:  "16:00～23:00",
		Budget:    4000,
		Img:       "https://rimage.gnst.jp/rest/img/111111111111/1111.jpg",
		Latitude:  11.111111,
		Longitude: 11.111111,
		URL:       "https://r.gnavi.co.jp/111111111111/?ak=bbbbbbbb",
		CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	},
}

var favorites = []models.Favorite{
	{
		UserID:    1,
		ShopID:    1,
		CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	},
	{
		UserID:    2,
		ShopID:    1,
		CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
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
