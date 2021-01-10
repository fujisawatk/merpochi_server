package repository

import "merpochi_server/domain/models"

// UserRepository userPersistenceの抽象依存
type UserRepository interface {
	Save(models.User) (models.User, error)
	FindByID(uint32) (models.User, error)
	Update(uint32, models.User) (int64, error)
	Delete(uint32) (int64, error)
	FindBookmarkedShops(uint32) (*[]models.Shop, error)
	FindFavoritedShops(uint32) (*[]models.Shop, error)
	FindMyPosts(uint32) (*[]models.Post, error)
	FindCommentedPosts(uint32) (*[]models.Post, error)
}
