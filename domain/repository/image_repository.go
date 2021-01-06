package repository

import (
	"bytes"
	"merpochi_server/domain/models"
)

// ImageRepository imagePersistenceの抽象依存
type ImageRepository interface {
	Upload(string, *bytes.Buffer) error
	Create(models.Image) (models.Image, error)
}
