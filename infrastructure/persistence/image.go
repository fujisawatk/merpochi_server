package persistence

import (
	"fmt"
	"merpochi_server/config"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/util/channels"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jinzhu/gorm"
)

type imagePersistence struct {
	db *gorm.DB
}

// NewImagePersistence userPersistence構造体の宣言
func NewImagePersistence(db *gorm.DB) repository.ImageRepository {
	return &imagePersistence{db}
}

// Upload ユーザー画像をAmazon S3へアップロード
func (ip *imagePersistence) Upload(img *models.Image) error {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		creds := credentials.NewStaticCredentials(config.AWSID, config.AWSKEY, "")
		sess, err := session.NewSession(&aws.Config{
			Credentials: creds,
			Region:      aws.String("ap-northeast-1")},
		)
		if err != nil {
			ch <- false
			return
		}

		uploader := s3manager.NewUploader(sess)
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String("merpochi-users-image"),
			Key:    aws.String(img.Name),
			Body:   img.Buf,
		})
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return nil
	}
	return err
}

func (ip *imagePersistence) Create(img *models.Image) (*models.Image, error) {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = ip.db.Model(&models.Image{}).Create(&img).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return img, nil
	}
	fmt.Println(img)
	return &models.Image{}, err
}
