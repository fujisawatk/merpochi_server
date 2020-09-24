package repository

import "merpochi_server/domain/models"

// ShopRepository shopPersistenceの抽象依存
type ShopRepository interface {
	FindCommentsCount(uint32) (uint32, error)
	FindFavoritesCount(uint32) (uint32, error)
	Save(models.Shop) (models.Shop, error)
	Search(string) (models.Shop, error)
	FindCommentedShops(uint32) ([]models.Shop, error)
	FindFavoritedShops(uint32) ([]models.Shop, error)
}
