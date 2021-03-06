package persistence

import (
	"merpochi_server/domain/models"
	"reflect"
	"testing"
	"time"
)

func TestFavorite_Save(t *testing.T) {
	tests := []struct {
		name    string
		args    models.Favorite
		want    models.Favorite
		wantErr bool
	}{
		{
			name: "指定した店舗IDに紐付くお気に入り情報を保存出来ること",
			args: models.Favorite{
				UserID:    1,
				ShopID:    2,
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: models.Favorite{
				ID:        3,
				UserID:    1,
				ShopID:    2,
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			wantErr: false,
		},
		{
			name: "お気に入り登録済の場合、エラーを返すこと",
			args: models.Favorite{
				UserID:    1,
				ShopID:    1,
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want:    models.Favorite{},
			wantErr: true,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp := NewFavoritePersistence(tx)
			got, err := fp.Save(tt.args)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("favoritePersistence.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("favoritePersistence.Save() = %v, want %v", got, tt.want)
			}
		})
	}
	tx.Rollback()
}

func TestFavorite_Delete(t *testing.T) {
	type args struct {
		sid uint32
		uid uint32
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "お気に入り情報を削除出来ること",
			args: args{
				sid: 1,
				uid: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "一致するお気に入り情報が存在しない場合、エラーが返ること",
			args: args{
				sid: 2,
				uid: 1,
			},
			want:    0,
			wantErr: true,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp := NewFavoritePersistence(tx)
			got, err := fp.Delete(tt.args.sid, tt.args.uid)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("favoritePersistence.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("favoritePersistence.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
	tx.Rollback()
}

// func TestFavorite_FindFavoriteUser(t *testing.T) {
// 	type args struct {
// 		uid uint32
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    models.User
// 		wantErr bool
// 	}{
// 		{
// 			name: "お気に入りしたユーザー情報を取得出来ること",
// 			args: args{
// 				uid: 1,
// 			},
// 			want: models.User{
// 				Email: "miku@email.com",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "指定のユーザーIDが存在しない場合、エラーが返ること",
// 			args: args{
// 				uid: 10,
// 			},
// 			want:    models.User{},
// 			wantErr: true,
// 		},
// 	}
// 	tx := db.Begin()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			fp := NewFavoritePersistence(tx)
// 			got, err := fp.FindFavoriteUser(tt.args.uid)
// 			// 予期しないエラーの場合
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("favoritePersistence.FindFavoriteUser() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			// 返り値が期待しない値の場合
// 			if !reflect.DeepEqual(got.Email, tt.want.Email) {
// 				t.Errorf("favoritePersistence.FindFavoriteUser() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// 	tx.Rollback()
// }
