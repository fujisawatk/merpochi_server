package usecase

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/util/security"
	"mime/multipart"

	"github.com/nfnt/resize"
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
	bf := new(bytes.Buffer)
	defer bf.Reset()

	_, err := bf.ReadFrom(file)
	if err != nil {
		return models.Image{}, err
	}

	filename, err := ResizeImage(bf)
	if err != nil {
		return models.Image{}, err
	}

	err = iu.imageRepository.Upload(filename, bf)
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

// ResizeImage 画像の整形
func ResizeImage(bf *bytes.Buffer) (string, error) {
	var filename string

	img, t, err := image.Decode(bf)
	if err != nil {
		return "", err
	}

	m := resize.Resize(300, 0, img, resize.Lanczos3)
	switch t {
	case "jpeg":
		filename = security.RandomString(20) + ".jpg"

		err = jpeg.Encode(bf, m, nil)
		if err != nil {
			return "", err
		}
	case "png":
		filename = security.RandomString(20) + ".png"

		err = png.Encode(bf, m)
		if err != nil {
			return "", err
		}
	case "gif":
		filename = security.RandomString(20) + ".gif"

		err = gif.Encode(bf, m, nil)
		if err != nil {
			return "", err
		}
	}
	return filename, nil
}
