package repository

import "merpochi_server/domain/models"

// ShopRepository shopPersistenceの抽象依存
type ShopRepository interface {
	FindPostsCount(uint32) uint32
	FindFavoritesCount(uint32) uint32
	Save(models.Shop) (models.Shop, error)
	Search(string) (*models.Shop, error)
	FindCommentedShops(uint32) ([]models.Shop, error)
	FindFavoritedShops(uint32) ([]models.Shop, error)
}
