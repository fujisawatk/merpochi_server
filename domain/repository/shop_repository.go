package repository

import "merpochi_server/domain/models"

// ShopRepository shopPersistenceの抽象依存
type ShopRepository interface {
	FindAll(uint32) (uint32, error)
	Save(models.Shop) (models.Shop, error)
	FindByID(uint32) ([]models.Comment, error)
	SearchShop(string) (models.Shop, error)
}
