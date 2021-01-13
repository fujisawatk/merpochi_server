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
	FindAll(uint32, uint32, uint32) (*[]models.Image, error)
	UploadS3(*models.Image, string) error
	DownloadS3(*models.Image, string) error
	DeleteS3(*models.Image, string) error
}
