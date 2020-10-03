package persistence

import (
	"merpochi_server/domain/models"
	"reflect"
	"testing"
)

func TestFindCommentsCount(t *testing.T) {
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
			want:    1,
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewShopPersistence(db)
			got := sp.FindCommentsCount(tt.args.sid)
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shopPersistence.FindCommentsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindFavoritesCount(t *testing.T) {
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
			want:    1,
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewShopPersistence(db)
			got := sp.FindFavoritesCount(tt.args.sid)
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shopPersistence.FindFavoritesCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShopSave(t *testing.T) {
	tests := []struct {
		name    string
		args    models.Shop
		want    models.Shop
		wantErr bool
	}{
		{
			name: "店舗情報が登録出来ること",
			args: models.Shop{
				Code:      "bbbb111",
				Name:      "イタリアンショップ",
				Category:  "イタリアン",
				Opentime:  "17:00～23:00",
				Budget:    2000,
				Img:       "https://rimage.gnst.jp/rest/img/111111110000/1111.jpg",
				Latitude:  11.111111,
				Longitude: 11.111111,
				URL:       "https://r.gnavi.co.jp/111111110000/?ak=bbbbbbbb",
			},
			want: models.Shop{
				Code:      "bbbb111",
				Name:      "イタリアンショップ",
				Category:  "イタリアン",
				Opentime:  "17:00～23:00",
				Budget:    2000,
				Img:       "https://rimage.gnst.jp/rest/img/111111110000/1111.jpg",
				Latitude:  11.111111,
				Longitude: 11.111111,
				URL:       "https://r.gnavi.co.jp/111111110000/?ak=bbbbbbbb",
			},
			wantErr: false,
		},
		{
			name: "特定の値が空でも、登録出来ること(not null)",
			args: models.Shop{
				Code:      "cccc222",
				Name:      "",
				Category:  "",
				Opentime:  "",
				Budget:    0,
				Img:       "",
				Latitude:  0,
				Longitude: 0,
				URL:       "",
			},
			want: models.Shop{
				Code:      "cccc222",
				Name:      "",
				Category:  "",
				Opentime:  "",
				Budget:    0,
				Img:       "",
				Latitude:  0,
				Longitude: 0,
				URL:       "",
			},
			wantErr: false,
		},
		{
			name: "店舗コードが重複していたら、登録出来ないこと(unique)",
			args: models.Shop{
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
			want:    models.Shop{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewShopPersistence(db)
			got, err := sp.Save(tt.args)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("shopPersistence.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got.Code, tt.want.Code) {
				t.Errorf("shopPersistence.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShopSearch(t *testing.T) {
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
			wantErr: false,
		},
		{
			name: "指定した店舗コードが未登録の場合、エラーを返すこと",
			args: args{
				code: "dddd444",
			},
			want:    models.Shop{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := NewShopPersistence(db)
			got, err := sp.Search(tt.args.code)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("shopPersistence.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("shopPersistence.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
