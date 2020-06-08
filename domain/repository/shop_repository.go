package repository

import "merpochi_server/domain/models"

// UserRepository userPersistenceの抽象依存
type ShopRepository interface {
	FindAll() ([]models.Shop, error)
}
