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
	UploadImage(uint32, multipart.File) (*models.Image, error)
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

func (iu imageUsecase) UploadImage(uid uint32, file multipart.File) (*models.Image, error) {
	img := &models.Image{
		UserID: uid,
		Buf:    &bytes.Buffer{},
	}

	_, err := img.Buf.ReadFrom(file)
	if err != nil {
		return &models.Image{}, err
	}

	err = ResizeImage(img)
	if err != nil {
		return &models.Image{}, err
	}

	err = iu.imageRepository.Upload(img)
	if err != nil {
		return &models.Image{}, err
	}

	img, err = iu.imageRepository.Create(img)
	if err != nil {
		return &models.Image{}, err
	}

	return img, nil
}

// ResizeImage 画像の整形
func ResizeImage(i *models.Image) error {
	img, t, err := image.Decode(i.Buf)
	if err != nil {
		return err
	}

	m := resize.Resize(300, 0, img, resize.Lanczos3)
	switch t {
	case "jpeg":
		i.Name = security.RandomString(20) + ".jpg"

		err = jpeg.Encode(i.Buf, m, nil)
		if err != nil {
			return err
		}
	case "png":
		i.Name = security.RandomString(20) + ".png"

		err = png.Encode(i.Buf, m)
		if err != nil {
			return err
		}
	case "gif":
		i.Name = security.RandomString(20) + ".gif"

		err = gif.Encode(i.Buf, m, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
