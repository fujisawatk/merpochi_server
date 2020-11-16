package persistence

import (
	"merpochi_server/domain/models"
	"reflect"
	"testing"
	"time"
)

func TestFavoriteFindAll(t *testing.T) {
	type args struct {
		sid uint32
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Favorite
		wantErr bool
	}{
		{
			name: "指定した店舗IDに紐付くお気に入り情報を取得出来ること",
			args: args{
				sid: 1,
			},
			want: []models.Favorite{
				{
					ID:        1,
					UserID:    1,
					ShopID:    1,
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				},
				{
					ID:        2,
					UserID:    2,
					ShopID:    1,
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				},
			},
			wantErr: false,
		},
		{
			name: "指定した店舗IDに紐付くお気に入り情報が無い場合、空の値を返す",
			args: args{
				sid: 2,
			},
			want:    []models.Favorite{},
			wantErr: false,
		},
	}
	tx := db.Begin()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp := NewFavoritePersistence(tx)
			got, err := fp.FindAll(tt.args.sid)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("commentPersistence.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commentPersistence.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
	tx.Rollback()
}
