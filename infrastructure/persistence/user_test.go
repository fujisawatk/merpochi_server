package persistence

import (
	"merpochi_server/domain/models"
	"reflect"
	"strings"
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
			name: "ニックネームが20文字以内の場合、登録出来ること",
			args: models.User{
				Nickname: strings.Repeat("a", 20),
				Email:    "test1@email.com",
				Password: "testpassword",
			},
			want: models.User{
				Email: "test1@email.com",
			},
			wantErr: false,
		},
		{
			name: "ニックネームが21文字以上の場合、登録出来ないこと",
			args: models.User{
				Nickname: strings.Repeat("a", 21),
				Email:    "test2@email.com",
				Password: "testpassword",
			},
			want:    models.User{},
			wantErr: true,
		},
		{
			name: "メールアドレスが100文字以内の場合、登録出来ること",
			args: models.User{
				Nickname: "testname",
				Email:    strings.Repeat("a", 90) + "@email.com", // 100文字
				Password: "testpassword",
			},
			want: models.User{
				Email: strings.Repeat("a", 90) + "@email.com", // 100文字
			},
			wantErr: false,
		},
		{
			name: "メールアドレスが101文字以上の場合、登録出来ないこと",
			args: models.User{
				Nickname: "testname",
				Email:    strings.Repeat("b", 91) + "@email.com", // 101文字
				Password: "testpassword",
			},
			want:    models.User{},
			wantErr: true,
		},
		{
			name: "メールアドレスが重複している場合、登録出来ないこと",
			args: models.User{
				Nickname: "testname",
				Email:    "miku@email.com",
				Password: "testpassword",
			},
			want:    models.User{},
			wantErr: true,
		},
		{
			name: "パスワードが40文字以内の場合、登録出来ること",
			args: models.User{
				Nickname: "testname",
				Email:    "test3@email.com",
				Password: strings.Repeat("a", 40),
			},
			want: models.User{
				Email: "test3@email.com",
			},
			wantErr: false,
		},
		{
			name: "パスワードが41文字以上の場合、登録出来ないこと",
			args: models.User{
				Nickname: "testname",
				Email:    "fuji4@email.com",
				Password: strings.Repeat("b", 41),
			},
			want:    models.User{},
			wantErr: true,
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
			name: "指定したユーザー情報がない場合、エラーを返すこと",
			args: args{
				uid: 10,
			},
			want:    models.User{},
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
			name: "変更するニックネームが20文字以内の場合、更新出来ること",
			args: args{
				uid: 1,
				user: models.User{
					Nickname: strings.Repeat("b", 20),
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "変更するニックネームが21文字以上の場合、更新出来ないこと",
			args: args{
				uid: 2,
				user: models.User{
					Nickname: strings.Repeat("b", 21),
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "変更するメールアドレスが100文字以内の場合、更新出来ること",
			args: args{
				uid: 1,
				user: models.User{
					Email: strings.Repeat("a", 90) + "@email.com", // 100文字
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "変更するメールアドレスが101文字以上の場合、更新出来ないこと",
			args: args{
				uid: 2,
				user: models.User{
					Email: strings.Repeat("a", 91) + "@email.com", // 101文字
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "変更するメールアドレスが重複している場合、更新出来ないこと",
			args: args{
				uid: 3,
				user: models.User{
					Email: "taka@email.com",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "指定したユーザーIDが存在しない場合、更新出来ないこと",
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
