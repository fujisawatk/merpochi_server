package repository

import "merpochi_server/domain/models"

// CommentRepository commentPersistenceの抽象依存
type CommentRepository interface {
	FindAll(uint32) (*[]models.Comment, error)
	Save(*models.Comment) error
	Update(uint32, *models.Comment) (int64, error)
	Delete(uint32) error
	FindByUserID(uint32) (*models.User, error)
}
