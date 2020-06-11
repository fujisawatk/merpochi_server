package repository

import "merpochi_server/domain/models"

// CommentRepository commentPersistenceの抽象依存
type CommentRepository interface {
	Save(models.Comment) (models.Comment, error)
}
