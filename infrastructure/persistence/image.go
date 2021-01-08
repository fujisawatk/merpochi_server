package persistence

import (
	"bytes"
	"errors"
	"fmt"
	"merpochi_server/config"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/util/channels"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

// Search 重複確認（ユーザー画像は一意性の必要があるため）
func (ip *imagePersistence) Search(uid uint32) error {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		// ユーザー画像は一意性であるため
		result := ip.db.Model(&models.Image{}).Where("user_id = ? AND shop_id = ?", uid, 0).Take(&models.Image{})
		if result.RowsAffected > 0 {
			err = errors.New("user image is already registered")
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

// FindByID 指定ユーザーの画像情報を取得
func (ip *imagePersistence) FindByID(uid uint32) (*models.Image, error) {
	var err error
	img := &models.Image{
		Buf: &bytes.Buffer{},
	}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = ip.db.Model(&models.Image{}).Where("user_id = ? AND shop_id = ?", uid, 0).Take(img).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return img, nil
	}
	return &models.Image{}, err
}

func (ip *imagePersistence) Update(img *models.Image) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		rs = ip.db.Model(&models.Image{}).Where("id = ?", img.ID).Take(&models.Image{}).UpdateColumns(
			map[string]interface{}{
				"name":       img.Name,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

// Upload ユーザー画像をAmazon S3へアップロード
func (ip *imagePersistence) UploadS3(img *models.Image) error {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		sess, err := SetCredentialsForAWS()
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

// Download ユーザー画像をAmazon S3からダウンロード
func (ip *imagePersistence) DownloadS3(img *models.Image) error {
	var err error
	done := make(chan bool)
	fmt.Println(img.Name)
	go func(ch chan<- bool) {
		defer close(ch)

		sess, err := SetCredentialsForAWS()
		if err != nil {
			ch <- false
			return
		}
		file, err := os.Create("tmp/" + img.Name)
		if err != nil {
			ch <- false
			return
		}
		defer os.Remove("tmp/" + img.Name)

		downloader := s3manager.NewDownloader(sess)
		_, err = downloader.Download(file,
			&s3.GetObjectInput{
				Bucket: aws.String("merpochi-users-image"),
				Key:    aws.String(img.Name),
			})
		if err != nil {
			ch <- false
			return
		}
		_, err = img.Buf.ReadFrom(file)
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

// DeleteS3 Amazon S3へアップロードしたユーザー画像削除
func (ip *imagePersistence) DeleteS3(img *models.Image) error {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		sess, err := SetCredentialsForAWS()
		if err != nil {
			ch <- false
			return
		}
		svc := s3.New(sess)
		// 指定した画像の削除指示
		_, err = svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String("merpochi-users-image"),
			Key:    aws.String(img.Name),
		})
		if err != nil {
			ch <- false
			return
		}
		// 画像が正常に削除されたか確認
		err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
			Bucket: aws.String("merpochi-users-image"),
			Key:    aws.String(img.Name),
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

// SetCredentialsForAWS AWSの認証情報を設定
func SetCredentialsForAWS() (*session.Session, error) {
	creds := credentials.NewStaticCredentials(config.AWSID, config.AWSKEY, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-northeast-1")},
	)
	return sess, err
}
