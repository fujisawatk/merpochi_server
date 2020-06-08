package validations

import (
	"errors"
	"merpochi_server/domain/models"

	"github.com/badoux/checkmail"
)

// UserCreateValidate ユーザー新規登録時のサーバー側バリデーション処理
func UserCreateValidate(user *models.User) error {
	if user.Nickname == "" {
		return errors.New("required nickname")
	}
	if len(user.Nickname) > 20 {
		return errors.New("nickname is 20 characters or less")
	}

	if user.Email == "" {
		return errors.New("required email")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email")
	}

	if user.Password == "" {
		return errors.New("required password")
	}
	if len(user.Password) < 5 {
		return errors.New("password is 5 characters or more")
	}
	if len(user.Password) > 20 {
		return errors.New("password is 20 characters or less")
	}
	return nil
}

// UserUpdateValidate ユーザー情報更新時のサーバー側バリデーション処理
func UserUpdateValidate(user *models.User) error {
	if user.Nickname == "" {
		return errors.New("required nickname")
	}
	if len(user.Nickname) > 20 {
		return errors.New("nickname is 20 characters or less")
	}

	if user.Email == "" {
		return errors.New("required email")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email")
	}
	return nil
}
