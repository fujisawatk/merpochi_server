package validations

import (
	"errors"
	"merpochi_server/domain/models"
)

// CommentValidate コメント登録・編集時のサーバー側バリデーション処理
func CommentValidate(comment *models.Comment) error {
	if comment.Text == "" {
		return errors.New("required comment")
	}
	if len(comment.Text) > 255 {
		return errors.New("comment is 255 characters or less")
	}
	return nil
}
