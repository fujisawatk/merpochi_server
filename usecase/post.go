package usecase

import (
	"bytes"
	"encoding/base64"
	"fmt"
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
	GetPosts(uint32) ([]postsGetResponse, error)
	GetPost(uint32, uint32) (postGetResponse, error)
	UpdatePost(uint32, uint32, string) (int64, error)
	DeletePost(uint32) error
}

type postUsecase struct {
	postRepository    repository.PostRepository
	commentRepository repository.CommentRepository // has many
	imageRepository   repository.ImageRepository   // has many
}

// NewPostUsecase Postデータに関するUsecaseを生成
func NewPostUsecase(pr repository.PostRepository, cr repository.CommentRepository, ir repository.ImageRepository) PostUsecase {
	return &postUsecase{
		postRepository:    pr,
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
			fmt.Println(buf)
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
			img, err = pu.imageRepository.Create(img)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (pu *postUsecase) GetPosts(sid uint32) ([]postsGetResponse, error) {
	var responses []postsGetResponse

	posts, err := pu.postRepository.FindAll(sid)
	if err != nil {
		return []postsGetResponse{}, err
	}
	// 投稿が存在する場合
	if len(*posts) > 0 {
		// 投稿したユーザー情報を取得
		for i := 0; i < len(*posts); i++ {
			postedUser, imgURI, time, err := pu.GetUserData((*posts)[i].UserID, (*posts)[i].CreatedAt, (*posts)[i].UpdatedAt)
			if err != nil {
				return []postsGetResponse{}, err
			}
			// コメント数取得
			commentsCount := pu.postRepository.FindCommentsCount((*posts)[i].ID)

			res := postsGetResponse{
				ID:            (*posts)[i].ID,
				Text:          (*posts)[i].Text,
				Rating:        (*posts)[i].Rating,
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

func (pu *postUsecase) GetPost(sid, pid uint32) (postGetResponse, error) {
	post, err := pu.postRepository.FindByID(sid, pid)
	if err != nil {
		return postGetResponse{}, err
	}
	// 投稿したユーザー情報を取得
	postedUser, imgURI, time, err := pu.GetUserData((*post).UserID, (*post).CreatedAt, (*post).UpdatedAt)
	if err != nil {
		return postGetResponse{}, err
	}

	// 指定の投稿に紐付くコメントを全件取得
	comments, err := pu.commentRepository.FindAll(pid)
	if err != nil {
		return postGetResponse{}, err
	}

	var commentsData []commentData
	// コメントが存在する場合
	if len(*comments) > 0 {
		// 投稿にコメントしたユーザー情報を取得
		for i := 0; i < len(*comments); i++ {
			commentedUser, imgURI, time, err := pu.GetUserData((*comments)[i].UserID, (*comments)[i].CreatedAt, (*comments)[i].UpdatedAt)
			if err != nil {
				return postGetResponse{}, err
			}
			data := commentData{
				ID:           (*comments)[i].ID,
				Text:         (*comments)[i].Text,
				UserID:       (*comments)[i].UserID,
				UserNickname: commentedUser,
				UserImage:    imgURI,
				Time:         time,
			}
			commentsData = append(commentsData, data)
		}
	}

	res := postGetResponse{
		ID:           (*post).ID,
		Text:         (*post).Text,
		Rating:       (*post).Rating,
		UserID:       (*post).UserID,
		UserNickname: postedUser,
		UserImage:    imgURI,
		Comments:     commentsData,
		Time:         time,
	}
	return res, nil
}

func (pu *postUsecase) UpdatePost(pid, rating uint32, text string) (int64, error) {
	post := &models.Post{
		ID:     pid,
		Text:   text,
		Rating: rating,
	}

	err := validations.PostValidate(post)
	if err != nil {
		return 0, err
	}

	rows, err := pu.postRepository.Update(post)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (pu *postUsecase) DeletePost(pid uint32) error {
	err := pu.postRepository.Delete(pid)
	if err != nil {
		return err
	}
	return nil
}

// ユーザー情報取得〜整形まで
func (pu *postUsecase) GetUserData(uid uint32, createdAt, updatedAt time.Time) (string, string, string, error) {
	user, err := pu.postRepository.FindByUserID(uid)
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
	img, err := pu.imageRepository.FindByID(uid)
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

type postsGetResponse struct {
	ID            uint32 `json:"id"`
	Text          string `json:"text"`
	Rating        uint32 `json:"rating"`
	UserID        uint32 `json:"user_id"`
	UserNickname  string `json:"user_nickname"`
	UserImage     string `json:"user_image"`
	CommentsCount uint32 `json:"comments_count"`
	Time          string `json:"time"`
}

type postGetResponse struct {
	ID           uint32        `json:"id"`
	Text         string        `json:"text"`
	Rating       uint32        `json:"rating"`
	UserID       uint32        `json:"user_id"`
	UserNickname string        `json:"user_nickname"`
	UserImage    string        `json:"user_image"`
	Comments     []commentData `json:"comments"`
	Time         string        `json:"time"`
}

type commentData struct {
	ID           uint32 `json:"id"`
	Text         string `json:"text"`
	UserID       uint32 `json:"user_id"`
	UserNickname string `json:"user_nickname"`
	UserImage    string `json:"user_image"`
	Time         string `json:"time"`
}
