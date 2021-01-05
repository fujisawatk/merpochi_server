package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/util/security"
	"mime/multipart"
	"net/http"
)

// ImageUsecase Imageに対するUsecaseのインターフェイス
type ImageUsecase interface {
	UploadImage(uint32, multipart.File) (models.Image, error)
}

type imageUsecase struct {
	imageRepository repository.ImageRepository
}

// NewImageUsecase Imageデータに関するUsecaseを生成
func NewImageUsecase(ir repository.ImageRepository) ImageUsecase {
	return &imageUsecase{
		imageRepository: ir,
	}
}

func (iu imageUsecase) UploadImage(uid uint32, file multipart.File) (models.Image, error) {
	var filename string

	fileHeader := make([]byte, 512)

	_, err := file.Read(fileHeader)
	if err != nil {
		return models.Image{}, err
	}

	switch http.DetectContentType(fileHeader) {
	case "image/jpeg":
		filename = security.RandomString(20) + ".jpeg"
	}

	err = iu.imageRepository.Upload(filename, file)
	if err != nil {
		return models.Image{}, err
	}

	img := models.Image{
		Name:   filename,
		UserID: uid,
	}

	img, err = iu.imageRepository.Create(img)
	if err != nil {
		return models.Image{}, err
	}
	return img, nil
}
