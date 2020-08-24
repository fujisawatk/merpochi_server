package repository

import "merpochi_server/domain/models"

// FavoriteRepository favoritePersistenceの抽象依存
type FavoriteRepository interface {
	Save(models.Favorite) (models.Favorite, error)
	Delete(uint32, uint32) (int64, error)
	Search(uint32) (uint32, error)
	FindFavoriteUser(uint32) (models.User, error)
}
