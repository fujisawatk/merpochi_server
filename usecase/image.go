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
	"strconv"

	"github.com/nfnt/resize"
)

// ImageUsecase Imageに対するUsecaseのインターフェイス
type ImageUsecase interface {
	CreateImage(uint32, multipart.File) (*models.Image, error)
	GetImage(uint32) (string, error)
	UpdateImage(uint32, multipart.File) (int64, error)
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

func (iu imageUsecase) CreateImage(uid uint32, file multipart.File) (*models.Image, error) {
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

	err = iu.imageRepository.Search(uid)
	if err != nil {
		return &models.Image{}, err
	}

	err = iu.imageRepository.UploadS3(img, "merpochi-users-image")
	if err != nil {
		return &models.Image{}, err
	}

	img, err = iu.imageRepository.Save(img)
	if err != nil {
		return &models.Image{}, err
	}

	return img, nil
}

func (iu imageUsecase) GetImage(uid uint32) (string, error) {
	img, err := iu.imageRepository.FindByUserID(uid)
	if err != nil {
		return "", err
	}

	err = iu.imageRepository.DownloadS3(img, "merpochi-users-image")
	if err != nil {
		return "", err
	}

	uri, err := security.Base64EncodeToString(img.Buf)
	if err != nil {
		return "", err
	}
	return uri, nil
}

func (iu imageUsecase) UpdateImage(uid uint32, file multipart.File) (int64, error) {
	img, err := iu.imageRepository.FindByUserID(uid)
	if err != nil {
		return 0, err
	}
	// 旧画像の削除処理
	err = iu.imageRepository.DeleteS3(img, "merpochi-users-image")
	if err != nil {
		return 0, err
	}

	_, err = img.Buf.ReadFrom(file)
	if err != nil {
		return 0, err
	}

	err = ResizeImage(img)
	if err != nil {
		return 0, err
	}

	err = iu.imageRepository.UploadS3(img, "merpochi-users-image")
	if err != nil {
		return 0, err
	}

	rows, err := iu.imageRepository.Update(img)
	if err != nil {
		return 0, err
	}

	return rows, nil
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
		i.Name = strconv.Itoa(int(i.UserID)) + "-profile-image.jpg"

		err = jpeg.Encode(i.Buf, m, nil)
		if err != nil {
			return err
		}
	case "png":
		i.Name = strconv.Itoa(int(i.UserID)) + "-profile-image.png"

		err = png.Encode(i.Buf, m)
		if err != nil {
			return err
		}
	case "gif":
		i.Name = strconv.Itoa(int(i.UserID)) + "-profile-image.gif"
		err = gif.Encode(i.Buf, m, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

type imageData struct {
	ID  uint32 `json:"id"`
	URI string `json:"uri"`
}
