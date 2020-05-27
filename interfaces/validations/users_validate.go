package validations

import (
	"merpochi_server/domain/models"

	"github.com/go-playground/validator/v10"
)

// UserCreateValidate ユーザー新規登録時のサーバー側バリデーション処理
func UserCreateValidate(user *models.User) []string {
	var errorMessages []string

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			var errorMessage string
			fieldName := err.Field()

			switch fieldName {
			case "Nickname":
				errorMessage = "ユーザーネームが不正です"
			case "Email":
				errorMessage = "メールアドレスが不正です"
			case "Password":
				errorMessage = "パスワードが不正です"
			}
			errorMessages = append(errorMessages, errorMessage)
		}
		return errorMessages
	}
	return nil
}

// UserUpdateValidate ユーザー新規登録時のサーバー側バリデーション処理
func UserUpdateValidate(user *models.User) []string {
	var errorMessages []string

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			var errorMessage string
			fieldName := err.Field()

			switch fieldName {
			case "Nickname":
				errorMessage = "ユーザーネームが不正です"
			case "Email":
				errorMessage = "メールアドレスが不正です"
			}
			errorMessages = append(errorMessages, errorMessage)
		}
		return errorMessages
	}
	return nil
}
