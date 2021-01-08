package repository

import (
	"merpochi_server/domain/models"
)

// ImageRepository imagePersistenceの抽象依存
type ImageRepository interface {
	Create(*models.Image) (*models.Image, error)
	Search(uint32) error
	FindByID(uint32) (*models.Image, error)
	Update(*models.Image) (int64, error)
	Upload(*models.Image) error
	Download(*models.Image) error
	DeleteS3(*models.Image) error
}
