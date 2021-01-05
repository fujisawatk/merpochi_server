package repository

import (
	"merpochi_server/domain/models"
	"mime/multipart"
)

// ImageRepository imagePersistenceの抽象依存
type ImageRepository interface {
	Upload(string, multipart.File) error
	Create(models.Image) (models.Image, error)
}
