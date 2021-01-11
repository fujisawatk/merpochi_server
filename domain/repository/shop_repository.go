package repository

import "merpochi_server/domain/models"

// ShopRepository shopPersistenceの抽象依存
type ShopRepository interface {
	FindPostsCount(uint32) uint32
	FindFavoritesCount(uint32) uint32
	FindBookmarksCount(uint32) uint32
	FindBookmarkUser(uint32, uint32) bool
	Save(models.Shop) (models.Shop, error)
	FindByCode(string) (*models.Shop, error)
}
