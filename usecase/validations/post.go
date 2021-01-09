package validations

import (
	"errors"
	"merpochi_server/domain/models"
)

// PostCreateValidate 投稿新規登録時のサーバー側バリデーション処理
func PostCreateValidate(post *models.Post) error {
	if post.Text == "" {
		return errors.New("required text")
	}
	if len(post.Text) > 1000 {
		return errors.New("text is 1000 characters or less")
	}
	return nil
}
