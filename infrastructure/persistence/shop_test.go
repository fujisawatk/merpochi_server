package persistence

import (
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
