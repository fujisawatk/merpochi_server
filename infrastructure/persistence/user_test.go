package persistence

import (
	"merpochi_server/domain/models"
	"reflect"
	"testing"
)

func TestUser_Save(t *testing.T) {
	tests := []struct {
		name    string
		args    models.User
		want    models.User
		wantErr bool
	}{
		{
			name: "ユーザー情報の保存処理が正常に行われること",
			args: models.User{
				ID:       3,
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

func TestUser_FindByID(t *testing.T) {
	type args struct {
		uid uint32
	}
	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "指定したユーザー情報を取得出来ること",
			args: args{
				uid: 1,
			},
			want: models.User{
				Email: "miku@email.com",
			},
			wantErr: false,
		},
		{
			name: "指定したユーザー情報がない時にエラーを返すこと",
			args: args{
				uid: 10,
			},
			want: models.User{
				Email: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := NewUserPersistence(db)
			got, err := up.FindByID(tt.args.uid)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("userPersistence.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 返り値が期待しない値の場合
			if !reflect.DeepEqual(got.Email, tt.want.Email) {
				t.Errorf("userPersistence.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	type args struct {
		uid  uint32
		user models.User
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "指定したユーザー情報を更新出来ること",
			args: args{
				uid: 1,
				user: models.User{
					Email: "mikumiku@email.com",
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "指定したユーザー情報がない時は更新出来ないこと",
			args: args{
				uid: 10,
				user: models.User{
					Email: "mikumiku@email.com",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := NewUserPersistence(db)
			got, err := up.Update(tt.args.uid, tt.args.user)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("userPersistence.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 返り値が期待しない値の場合
			if got != tt.want {
				t.Errorf("userPersistence.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Delete(t *testing.T) {
	type args struct {
		uid uint32
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "指定したユーザー情報を削除出来ること",
			args: args{
				uid: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "指定したユーザー情報がない時は削除出来ないこと",
			args: args{
				uid: 10,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := NewUserPersistence(db)
			got, err := up.Delete(tt.args.uid)
			// 予期しないエラーの場合
			if (err != nil) != tt.wantErr {
				t.Errorf("userPersistence.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 返り値が期待しない値の場合
			if got != tt.want {
				t.Errorf("userPersistence.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SearchUser(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "指定したメールアドレスの重複が無いこと",
			args: args{
				email: "enako@email.com",
			},
			wantErr: false,
		},
		{
			name: "指定したメールアドレスが重複していること",
			args: args{
				email: "fuji@email.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := NewUserPersistence(db)
			if err := up.SearchUser(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("userPersistence.SearchUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
