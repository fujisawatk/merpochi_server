package persistence

import (
	"merpochi_server/domain/models"
	"reflect"
	"testing"
)

func TestSave(t *testing.T) {
	tests := []struct {
		name    string
		args    models.User
		want    models.User
		wantErr bool
	}{
		{
			name: "ユーザー情報の保存処理が正常に行われること",
			args: models.User{
				Nickname: "fujisawatk",
				Email:    "fuji@email.com",
				Password: "fujifuji0707",
			},
			want: models.User{
				Nickname: "fujisawatk",
				Email:    "fuji@email.com",
				Password: "fujifuji0707",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := NewUserPersistence(db)
			got, err := up.Save(tt.args)
			if (err != nil) != tt.wantErr {
				// 予期しないエラーの場合
				t.Errorf("userPersistence.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got.Email, tt.want.Email) {
				t.Errorf("userPersistence.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.User
		wantErr bool
	}{
		{
			name: "ユーザー情報を全件取得出来ること",
			want: []models.User{
				{
					Email: "miku@email.com",
				},
				{
					Email: "fuji@email.com",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := NewUserPersistence(db)
			got, err := up.FindAll()
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("userPersistence.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 返り値が期待しない値の場合
			if (got[0].Email != tt.want[0].Email) && (got[1].Email != tt.want[1].Email) {
				t.Errorf("userPersistence.FindAll() → %v, %v want → %v, %v",
					got[0].Email, got[1].Email, tt.want[0].Email, tt.want[1].Email)
			}
		})
	}
}
