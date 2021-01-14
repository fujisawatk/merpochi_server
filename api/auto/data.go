package auto

import (
	"merpochi_server/domain/models"
)

var users = []models.User{
	{
		ID:       1,
		Nickname: "miku",
		Email:    "miku@email.com",
		Password: "mikumiku",
		Genre:    "居酒屋",
	},
	{
		ID:       2,
		Nickname: "fuji",
		Email:    "fuji@email.com",
		Password: "fujifuji",
		Genre:    "焼き鳥",
	},
}

var shops = []models.Shop{
	{
		ID:        1,
		Code:      "gcrk447",
		Name:      "焼鳥屋 鳥貴族 鶴見西口店",
		Category:  "焼鳥",
		Opentime:  "17:00～24:00",
		Access:    "JR鶴見駅西口徒歩2分",
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
		Access:    "JR鶴見駅東口徒歩2分",
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
		UserID: 2,
		ShopID: 2,
	},
}

var comments = []models.Comment{
	{
		ID:     1,
		Text:   "コメントテストa00000",
		UserID: 1,
		PostID: 1,
	},
	{
		ID:     2,
		Text:   "コメントテストb11111",
		UserID: 2,
		PostID: 2,
	},
	{
		ID:     3,
		Text:   "コメントテストa00000-2",
		UserID: 1,
		PostID: 2,
	},
	{
		ID:     4,
		Text:   "コメントテストa00000-3",
		UserID: 1,
		PostID: 2,
	},
}

var images = []models.Image{
	{
		ID:     1,
		Name:   "1-profile-image.png",
		UserID: 1,
		ShopID: 0,
		PostID: 0,
	},
	{
		ID:     2,
		Name:   "2-profile-image.jpg",
		UserID: 2,
		ShopID: 0,
		PostID: 0,
	},
	{
		ID:     3,
		Name:   "shop-2-post-2-image-1.jpg",
		UserID: 2,
		ShopID: 2,
		PostID: 2,
	},
	{
		ID:     4,
		Name:   "shop-2-post-2-image-2.jpg",
		UserID: 2,
		ShopID: 2,
		PostID: 2,
	},
	{
		ID:     5,
		Name:   "shop-2-post-2-image-3.jpg",
		UserID: 2,
		ShopID: 2,
		PostID: 2,
	},
	{
		ID:     6,
		Name:   "shop-2-post-2-image-4.jpg",
		UserID: 2,
		ShopID: 2,
		PostID: 2,
	},
}

var posts = []models.Post{
	{
		ID:     1,
		Text:   "美味しかったです！",
		Rating: 5,
		ShopID: 1,
		UserID: 1,
	},
	{
		ID: 2,
		Text: "有名で、人気のお店(チェーン)ですね。地方では見かけませんが。\n" +
			"会社の同僚と何度か伺っています。店内はお客さんでいっぱいでも、注文して→スムーズに提供してくれます。\n" +
			"串のサイズ感はやや大きめで、食べごたえもあり、美味しいと思います。\n" +
			"串以外も(食べた中で)鳥の酢モノや鳥ご飯が美味しいと思います。",
		Rating: 4,
		ShopID: 2,
		UserID: 2,
	},
	{
		ID: 3,
		Text: "昔、名古屋にいた時に遅くまでやっているのと全てがほぼ均一で安かったので夜食代わりによく行っていた鳥貴族。\n" +
			"鶴見にもあったので軽く二軒目ということで入ってみた。いい意味でも悪い意味でも雰囲気も含め均一だった（笑）",
		Rating: 3,
		ShopID: 2,
		UserID: 1,
	},
	{
		ID:     4,
		Text:   "ぜひリピートしたいです！",
		Rating: 5,
		ShopID: 2,
		UserID: 1,
	},
}

var bookmarks = []models.Bookmark{
	{
		ID:     1,
		UserID: 1,
		ShopID: 1,
	},
	{
		ID:     2,
		UserID: 2,
		ShopID: 2,
	},
}
