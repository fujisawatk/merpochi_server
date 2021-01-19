package repository

import "merpochi_server/domain/models"

// ShopRepository shopPersistenceの抽象依存
type ShopRepository interface {
	Save(models.Shop) (models.Shop, error)
	FindByCode(string) (*models.Shop, error)
	FindByPostID(uint32) (*models.Shop, error)
	FindAllByUserIDJoinsBookmark(uint32) (*[]models.Shop, error)
	FindAllByUserIDJoinsFavorite(uint32) (*[]models.Shop, error)
}
