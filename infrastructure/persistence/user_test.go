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
				t.Errorf("userPersistence.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Email, tt.want.Email) {
				t.Errorf("userPersistence.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
