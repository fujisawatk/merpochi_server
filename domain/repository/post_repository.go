package repository

import "merpochi_server/domain/models"

// PostRepository postPersistenceの抽象依存
type PostRepository interface {
	Save(*models.Post) error
	FindAll(uint32) (*[]models.Post, error)
	FindByID(uint32, uint32) (*models.Post, error)
	Update(*models.Post) (int64, error)
	Delete(uint32) error
	FindByUserID(uint32) (*models.User, error)
	FindMyPosts(uint32) (*[]models.Post, error)
	FindCommentedPosts(uint32) (*[]models.Post, error)
	FindPostsCount(uint32) uint32
}
