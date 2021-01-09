package repository

import "merpochi_server/domain/models"

// PostRepository postPersistenceの抽象依存
type PostRepository interface {
	Save(*models.Post) error
	FindAll(uint32) (*[]models.Post, error)
}
