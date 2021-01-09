package validations

import (
	"errors"
	"merpochi_server/domain/models"
)

// PostValidate 投稿保存、更新時のサーバー側バリデーション処理
func PostValidate(post *models.Post) error {
	if post.Text == "" {
		return errors.New("required text")
	}
	if len(post.Text) > 1000 {
		return errors.New("text is 1000 characters or less")
	}
	return nil
}
