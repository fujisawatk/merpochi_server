package repository

import "merpochi_server/domain/models"

// FavoriteRepository favoritePersistenceの抽象依存
type FavoriteRepository interface {
	Save(models.Favorite) error
	Delete(uint32, uint32) error
	Search(uint32) (uint32, error)
}