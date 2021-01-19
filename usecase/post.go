package usecase

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/usecase/validations"
	"merpochi_server/util/security"
	"strconv"
	"strings"
	"time"

	"github.com/nfnt/resize"
)

// PostUsecase Postに対するUsecaseのインターフェイス
type PostUsecase interface {
	CreatePost([]string, string, uint32, uint32, uint32) error
	GetPosts(uint32) ([]getPostsResponse, error)
	UpdatePost([]string, string, uint32, uint32, uint32, uint32) error
	DeletePost(uint32) error
}

type postUsecase struct {
	postRepository    repository.PostRepository
	userRepository    repository.UserRepository    // belongs to
	commentRepository repository.CommentRepository // has many
	imageRepository   repository.ImageRepository   // has many
}

// NewPostUsecase Postデータに関するUsecaseを生成
func NewPostUsecase(
	pr repository.PostRepository,
	ur repository.UserRepository,
	cr repository.CommentRepository,
	ir repository.ImageRepository,
) PostUsecase {
	return &postUsecase{
		postRepository:    pr,
		userRepository:    ur,
		commentRepository: cr,
		imageRepository:   ir,
	}
}

func (pu *postUsecase) CreatePost(imgs []string, text string, rating, uid, sid uint32) error {
	post := &models.Post{
		Text:   text,
		Rating: rating,
		UserID: uid,
		ShopID: sid,
	}
	err := validations.PostValidate(post)
	if err != nil {
		return err
	}

	err = pu.postRepository.Save(post)
	if err != nil {
		return err
	}
	if len(imgs) > 0 {
		for i := 0; i < len(imgs); i++ {
			img := &models.Image{
				UserID: uid,
				ShopID: sid,
				PostID: (*post).ID,
				Buf:    &bytes.Buffer{},
			}
			// base64エンコード文字列を最初のコンマまでカット("data:image/png;base64,"部分がデコード時に不要のため )
			b64data := imgs[i][strings.IndexByte(imgs[i], ',')+1:]
			// 文字列をデコード
			data, err := base64.StdEncoding.DecodeString(b64data)
			if err != nil {
				return err
			}
			// バッファー生成
			buf := bytes.NewBuffer(data)
			_, err = img.Buf.ReadFrom(buf)
			if err != nil {
				return err
			}
			err = ResizePostImage(img, (i + 1))
			if err != nil {
				return err
			}
			err = pu.imageRepository.UploadS3(img, "merpochi-posts-image")
			if err != nil {
				return err
			}
			img, err = pu.imageRepository.Save(img)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (pu *postUsecase) GetPosts(sid uint32) ([]getPostsResponse, error) {
	var responses []getPostsResponse

	posts, err := pu.postRepository.FindAll(sid)
	if err != nil {
		return []getPostsResponse{}, err
	}
	// 投稿が存在する場合
	if len(*posts) > 0 {
		// 投稿したユーザー情報を取得
		for i := 0; i < len(*posts); i++ {
			postedUser, imgURI, time, err := pu.GetUserData((*posts)[i].UserID, (*posts)[i].CreatedAt, (*posts)[i].UpdatedAt)
			if err != nil {
				return []getPostsResponse{}, err
			}
			// コメント数取得
			commentsCount := pu.commentRepository.CountByPostID((*posts)[i].ID)

			imgs, err := pu.GetPostImage((*posts)[i].UserID, sid, (*posts)[i].ID)
			if err != nil {
				return []getPostsResponse{}, err
			}

			res := getPostsResponse{
				ID:            (*posts)[i].ID,
				Text:          (*posts)[i].Text,
				Rating:        (*posts)[i].Rating,
				Images:        imgs,
				UserID:        (*posts)[i].UserID,
				UserNickname:  postedUser,
				UserImage:     imgURI,
				CommentsCount: commentsCount,
				Time:          time,
			}
			responses = append(responses, res)
		}
	}
	return responses, nil
}

func (pu *postUsecase) UpdatePost(reImgs []string, text string, rating, uid, sid, pid uint32) error {
	post := &models.Post{
		ID:     pid,
		Text:   text,
		Rating: rating,
	}

	err := validations.PostValidate(post)
	if err != nil {
		return err
	}

	_, err = pu.postRepository.Update(post)
	if err != nil {
		return err
	}

	imgs, err := pu.imageRepository.FindAllByPostID(pid)
	if err != nil {
		return err
	}
	if len(*imgs) > 0 {
		// 旧画像の削除処理
		for _, i := range *imgs {
			tmp := &models.Image{
				Name: i.Name,
			}
			err = pu.imageRepository.DeleteS3(tmp, "merpochi-posts-image")
			if err != nil {
				return err
			}
			err = pu.imageRepository.DeleteByPostID(pid)
			if err != nil {
				return err
			}
		}
	}

	if len(reImgs) > 0 {
		for i := 0; i < len(reImgs); i++ {
			img := &models.Image{
				UserID: uid,
				ShopID: sid,
				PostID: (*post).ID,
				Buf:    &bytes.Buffer{},
			}
			// base64エンコード文字列を最初のコンマまでカット("data:image/png;base64,"部分がデコード時に不要のため )
			b64data := reImgs[i][strings.IndexByte(reImgs[i], ',')+1:]
			// 文字列をデコード
			data, err := base64.StdEncoding.DecodeString(b64data)
			if err != nil {
				return err
			}
			// バッファー生成
			buf := bytes.NewBuffer(data)
			_, err = img.Buf.ReadFrom(buf)
			if err != nil {
				return err
			}
			err = ResizePostImage(img, (i + 1))
			if err != nil {
				return err
			}
			err = pu.imageRepository.UploadS3(img, "merpochi-posts-image")
			if err != nil {
				return err
			}
			_, err = pu.imageRepository.Save(img)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (pu *postUsecase) DeletePost(pid uint32) error {
	// 投稿画像削除
	imgs, err := pu.imageRepository.FindAllByPostID(pid)
	if err != nil {
		return err
	}
	if len(*imgs) > 0 {
		for _, i := range *imgs {
			tmp := &models.Image{
				Name: i.Name,
			}
			err = pu.imageRepository.DeleteS3(tmp, "merpochi-posts-image")
			if err != nil {
				return err
			}
			err = pu.imageRepository.DeleteByPostID(pid)
			if err != nil {
				return err
			}
		}
	}

	err = pu.postRepository.Delete(pid)
	if err != nil {
		return err
	}
	return nil
}

// ユーザー情報取得〜整形まで
func (pu *postUsecase) GetUserData(uid uint32, createdAt, updatedAt time.Time) (string, string, string, error) {
	user, err := pu.userRepository.FindByID(uid)
	if err != nil {
		return "", "", "", err
	}

	imgURI, err := pu.GetUserImage(uid)
	if err != nil {
		return "", "", "", err
	}
	// 作成or編集時刻整形
	var time string
	format1 := "2006/01/02 15:04:05"
	if createdAt != updatedAt {
		time = "編集済 " + (updatedAt).Format(format1)
	} else {
		time = (createdAt).Format(format1)
	}
	return user.Nickname, imgURI, time, nil
}

// ユーザー画像取得〜base64エンコード文字列生成まで
func (pu *postUsecase) GetUserImage(uid uint32) (string, error) {
	img, err := pu.imageRepository.FindByUserID(uid)
	if err != nil {
		return "", err
	}

	err = pu.imageRepository.DownloadS3(img, "merpochi-users-image")
	if err != nil {
		return "", err
	}

	uri, err := security.Base64EncodeToString(img.Buf)
	if err != nil {
		return "", err
	}
	return uri, nil
}

// 投稿画像取得〜base64エンコード文字列生成まで
func (pu *postUsecase) GetPostImage(uid, sid, pid uint32) ([]imageData, error) {
	imgs, err := pu.imageRepository.FindAll(uid, sid, pid)
	if err != nil {
		return []imageData{}, err
	}
	var responses []imageData
	if len(*imgs) > 0 {
		for i := 0; i < len(*imgs); i++ {
			img := &models.Image{
				Name: (*imgs)[i].Name,
				Buf:  &bytes.Buffer{},
			}
			err = pu.imageRepository.DownloadS3(img, "merpochi-posts-image")
			if err != nil {
				return []imageData{}, err
			}

			uri, err := security.Base64EncodeToString((*img).Buf)
			if err != nil {
				return []imageData{}, err
			}
			res := imageData{
				ID:  (*imgs)[i].ID,
				URI: uri,
			}
			responses = append(responses, res)
		}
	}
	return responses, nil
}

// ResizePostImage 画像の整形
func ResizePostImage(i *models.Image, count int) error {
	img, t, err := image.Decode(i.Buf)
	if err != nil {
		return err
	}

	m := resize.Resize(300, 0, img, resize.Lanczos3)
	switch t {
	case "jpeg":
		i.Name = "shop-" + strconv.Itoa(int(i.ShopID)) +
			"-post-" + strconv.Itoa(int(i.PostID)) +
			"-image-" + strconv.Itoa(count) + ".jpg"

		err = jpeg.Encode(i.Buf, m, nil)
		if err != nil {
			return err
		}
	case "png":
		i.Name = "shop-" + strconv.Itoa(int(i.ShopID)) +
			"-post-" + strconv.Itoa(int(i.PostID)) +
			"-image-" + strconv.Itoa(count) + ".png"

		err = png.Encode(i.Buf, m)
		if err != nil {
			return err
		}
	case "gif":
		i.Name = "shop-" + strconv.Itoa(int(i.ShopID)) +
			"-post-" + strconv.Itoa(int(i.PostID)) +
			"-image-" + strconv.Itoa(count) + ".gif"

		err = gif.Encode(i.Buf, m, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

type getPostsResponse struct {
	ID            uint32      `json:"id"`
	Text          string      `json:"text"`
	Rating        uint32      `json:"rating"`
	Images        []imageData `json:"images"`
	UserID        uint32      `json:"user_id"`
	UserNickname  string      `json:"user_nickname"`
	UserImage     string      `json:"user_image"`
	CommentsCount uint32      `json:"comments_count"`
	Time          string      `json:"time"`
}

type postData struct {
	ID            uint32      `json:"id"`
	Text          string      `json:"text"`
	Rating        uint32      `json:"rating"`
	Images        []imageData `json:"images"`
	UserID        uint32      `json:"user_id"`
	UserNickname  string      `json:"user_nickname"`
	UserImage     string      `json:"user_image"`
	CommentsCount uint32      `json:"comments_count"`
	Time          string      `json:"time"`
}
