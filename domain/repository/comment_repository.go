package repository

import "merpochi_server/domain/models"

// CommentRepository commentPersistenceの抽象依存
type CommentRepository interface {
	FindAll(uint32) ([]models.Comment, error)
	Save(models.Comment) (models.Comment, error)
	Update(uint32, models.Comment) (int64, error)
	Delete(uint32) (int64, error)
	FindCommentUser(uint32) (models.User, error)
}
