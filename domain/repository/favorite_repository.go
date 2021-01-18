package repository

import "merpochi_server/domain/models"

// FavoriteRepository favoritePersistenceの抽象依存
type FavoriteRepository interface {
	FindAll(uint32) ([]models.Favorite, error)
	Save(models.Favorite) (models.Favorite, error)
	Delete(uint32, uint32) (int64, error)
	FindFavoriteUser(uint32) (models.User, error)
	CountByShopID(uint32) uint32
}
