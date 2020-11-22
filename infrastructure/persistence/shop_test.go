package persistence

import (
	"merpochi_server/domain/models"
	"reflect"
	"testing"
	"time"
)

func TestShop_FindCommentsCount(t *testing.T) {
	type args struct {
		sid uint32
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "指定した店舗情報に紐付くコメント数を取得出来ること",
			args: args{
				sid: 1,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "指定した店舗情報に紐付くコメントがない場合、'0'を返すこと",
			args: args{
				sid: 3,
			},
			want:    0,
			wantErr: false,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewShopPersistence(tx)
			got := sp.FindCommentsCount(tt.args.sid)
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shopPersistence.FindCommentsCount() = %v, want %v", got, tt.want)
			}
		})
	}
	tx.Rollback()
}

func TestShop_FindFavoritesCount(t *testing.T) {
	type args struct {
		sid uint32
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "指定した店舗情報に紐付くお気に入り数を取得出来ること",
			args: args{
				sid: 1,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "指定した店舗情報に紐付くお気に入りがない場合、'0'を返すこと",
			args: args{
				sid: 3,
			},
			want:    0,
			wantErr: false,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewShopPersistence(tx)
			got := sp.FindFavoritesCount(tt.args.sid)
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shopPersistence.FindFavoritesCount() = %v, want %v", got, tt.want)
			}
		})
	}
	tx.Rollback()
}

func TestShop_Save(t *testing.T) {
	tests := []struct {
		name    string
		args    models.Shop
		want    models.Shop
		wantErr bool
	}{
		{
			name: "店舗情報が登録出来ること",
			args: models.Shop{
				Code:      "cccc222",
				Name:      "イタリアンショップ",
				Category:  "イタリアン",
				Opentime:  "18:00～23:00",
				Budget:    2000,
				Img:       "https://rimage.gnst.jp/rest/img/222222222222/2222.jpg",
				Latitude:  22.222222,
				Longitude: 22.222222,
				URL:       "https://r.gnavi.co.jp/222222222222/?ak=cccccccc",
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: models.Shop{
				ID:        3,
				Code:      "cccc222",
				Name:      "イタリアンショップ",
				Category:  "イタリアン",
				Opentime:  "18:00～23:00",
				Budget:    2000,
				Img:       "https://rimage.gnst.jp/rest/img/222222222222/2222.jpg",
				Latitude:  22.222222,
				Longitude: 22.222222,
				URL:       "https://r.gnavi.co.jp/222222222222/?ak=cccccccc",
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			wantErr: false,
		},
		{
			name: "特定の値が空でも、登録出来ること(not null)",
			args: models.Shop{
				Code:      "dddd444",
				Name:      "",
				Category:  "",
				Opentime:  "",
				Budget:    0,
				Img:       "",
				Latitude:  0,
				Longitude: 0,
				URL:       "",
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: models.Shop{
				ID:        4,
				Code:      "dddd444",
				Name:      "",
				Category:  "",
				Opentime:  "",
				Budget:    0,
				Img:       "",
				Latitude:  0,
				Longitude: 0,
				URL:       "",
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			wantErr: false,
		},
		{
			name: "店舗コードが重複していたら、登録出来ないこと(unique)",
			args: models.Shop{
				Code: "aaaa000",
			},
			want:    models.Shop{},
			wantErr: true,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewShopPersistence(tx)
			got, err := sp.Save(tt.args)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("shopPersistence.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shopPersistence.Save() = %v, want %v", got, tt.want)
			}
		})
	}
	tx.Rollback()
}

func TestShop_Search(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Shop
		wantErr bool
	}{
		{
			name: "指定した店舗コードが登録されている場合、該当する店舗情報を取得出来ること",
			args: args{
				code: "aaaa000",
			},
			want: models.Shop{
				ID:        1,
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
			wantErr: false,
		},
		{
			name: "指定した店舗コードが未登録の場合、エラーを返すこと",
			args: args{
				code: "eeee555",
			},
			want:    models.Shop{},
			wantErr: true,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewShopPersistence(tx)
			got, err := sp.Search(tt.args.code)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("shopPersistence.Search() error = %v, wantErr %v", err, tt.wantErr)
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shopPersistence.Search() = %v, want %v", got, tt.want)
			}
		})
	}
	tx.Rollback()
}

func TestShop_FindCommentedShops(t *testing.T) {
	type args struct {
		uid uint32
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Shop
		wantErr bool
	}{
		{
			name: "ログインユーザーがコメントした店舗情報を取得出来ること",
			args: args{
				uid: 1,
			},
			want: []models.Shop{
				{
					ID:        1,
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
			},
			wantErr: false,
		},
		{
			name: "ログインユーザーがコメントした店舗がない場合、空の値を返す",
			args: args{
				uid: 3,
			},
			want:    []models.Shop{},
			wantErr: false,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewShopPersistence(tx)
			got, err := sp.FindCommentedShops(tt.args.uid)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("shopPersistence.FindCommentedShops() error = %v, wantErr %v", err, tt.wantErr)
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shopPersistence.FindCommentedShops() = %v, want %v", got, tt.want)
			}
		})
	}
	tx.Rollback()
}

func TestShop_FindFavoritedShops(t *testing.T) {
	type args struct {
		uid uint32
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Shop
		wantErr bool
	}{
		{
			name: "ログインユーザーがお気に入りした店舗情報を取得出来ること",
			args: args{
				uid: 1,
			},
			want: []models.Shop{
				{
					ID:        1,
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
			},
			wantErr: false,
		},
		{
			name: "ログインユーザーがお気に入りした店舗がない場合、空の値を返す",
			args: args{
				uid: 3,
			},
			want:    []models.Shop{},
			wantErr: false,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewShopPersistence(tx)
			got, err := sp.FindFavoritedShops(tt.args.uid)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("shopPersistence.FindFavoritedShops() error = %v, wantErr %v", err, tt.wantErr)
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shopPersistence.FindFaoviritedShops() = %v, want %v", got, tt.want)
			}
		})
	}
	tx.Rollback()
}
