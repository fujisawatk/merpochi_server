package repository

import "merpochi_server/domain/models"

// ShopRepository shopPersistenceの抽象依存
type ShopRepository interface {
	FindAll() ([]models.Shop, error)
	Save(models.Shop) (models.Shop, error)
	FindByID(uint32) ([]models.Comment, error)
}
