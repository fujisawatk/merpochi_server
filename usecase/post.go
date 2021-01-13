package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/usecase/validations"
	"merpochi_server/util/security"
)

// PostUsecase Postに対するUsecaseのインターフェイス
type PostUsecase interface {
	CreatePost(string, uint32, uint32, uint32) (*models.Post, error)
	GetPosts(uint32) ([]postsGetResponse, error)
	GetPost(uint32, uint32) (postGetResponse, error)
	UpdatePost(uint32, uint32, string) (int64, error)
	DeletePost(uint32) error
}

type postUsecase struct {
	postRepository    repository.PostRepository
	commentRepository repository.CommentRepository
	imageRepository   repository.ImageRepository
}

// NewPostUsecase Postデータに関するUsecaseを生成
func NewPostUsecase(pr repository.PostRepository, cr repository.CommentRepository, ir repository.ImageRepository) PostUsecase {
	return &postUsecase{
		postRepository:    pr,
		commentRepository: cr,
		imageRepository:   ir,
	}
}

func (pu *postUsecase) CreatePost(text string, rating, uid, sid uint32) (*models.Post, error) {
	post := &models.Post{
		Text:   text,
		Rating: rating,
		UserID: uid,
		ShopID: sid,
	}

	err := validations.PostValidate(post)
	if err != nil {
		return &models.Post{}, err
	}

	err = pu.postRepository.Save(post)
	if err != nil {
		return &models.Post{}, err
	}
	return post, nil
}

func (pu *postUsecase) GetPosts(sid uint32) ([]postsGetResponse, error) {
	var responses []postsGetResponse

	posts, err := pu.postRepository.FindAll(sid)
	if err != nil {
		return []postsGetResponse{}, err
	}
	// 投稿が存在する場合
	if len(*posts) > 0 {
		for i := 0; i < len(*posts); i++ {
			user, img, time, err := pu.GetOtherData((*posts)[i])
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
				UserNickname:  user.Nickname,
				UserImage:     img,
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
	// 指定の投稿に紐付くコメントを全件取得
	comments, err := pu.commentRepository.FindAll(pid)
	if err != nil {
		return postGetResponse{}, err
	}
	var commentsData []commentData
	if len(*comments) > 0 {
		// 取得した投稿にコメントしたユーザー情報を取得
		for i := 0; i < len(*comments); i++ {
			user, err := pu.commentRepository.FindByUserID((*comments)[i].UserID)
			if err != nil {
				return postGetResponse{}, err
			}
			// ユーザー画像取得
			imgURI, err := pu.GetImage((*comments)[i].UserID)
			if err != nil {
				return postGetResponse{}, err
			}

			// 投稿作成or編集時刻設定
			var time string
			format1 := "2006/01/02 15:04:05"
			if (*comments)[i].CreatedAt != (*comments)[i].UpdatedAt {
				time = "編集済 " + ((*comments)[i].UpdatedAt).Format(format1)
			} else {
				time = ((*comments)[i].CreatedAt).Format(format1)
			}
			data := commentData{
				ID:           (*comments)[i].ID,
				Text:         (*comments)[i].Text,
				UserID:       (*comments)[i].UserID,
				UserNickname: user.Nickname,
				UserImage:    imgURI,
				Time:         time,
			}
			commentsData = append(commentsData, data)
		}
	}
	user, img, time, err := pu.GetOtherData((*post))
	if err != nil {
		return postGetResponse{}, err
	}
	res := postGetResponse{
		ID:           (*post).ID,
		Text:         (*post).Text,
		Rating:       (*post).Rating,
		UserID:       (*post).UserID,
		UserNickname: user.Nickname,
		UserImage:    img,
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

// 投稿情報全件・一件取得機能のメソッド共有化
func (pu *postUsecase) GetOtherData(post models.Post) (*models.User, string, string, error) {
	// ユーザー値取得
	user, err := pu.postRepository.FindByUserID(post.UserID)
	if err != nil {
		return &models.User{}, "", "", err
	}
	// ユーザー画像取得
	imgURI, err := pu.GetImage(post.UserID)
	if err != nil {
		return &models.User{}, "", "", err
	}
	// 投稿作成or編集時刻設定
	var time string
	format1 := "2006/01/02 15:04:05"
	if post.CreatedAt != post.UpdatedAt {
		time = "編集済 " + (post.UpdatedAt).Format(format1)
	} else {
		time = (post.CreatedAt).Format(format1)
	}
	return user, imgURI, time, nil
}

// ユーザー画像取得〜base64エンコード文字列生成まで
func (pu *postUsecase) GetImage(uid uint32) (string, error) {
	img, err := pu.imageRepository.FindByID(uid)
	if err != nil {
		return "", err
	}

	err = pu.imageRepository.DownloadS3(img)
	if err != nil {
		return "", err
	}

	uri, err := security.Base64EncodeToString(img.Buf)
	if err != nil {
		return "", err
	}
	return uri, nil
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
