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
	GetPosts(uint32) ([]postsResponse, error)
	GetPost(uint32, uint32) (*models.Post, error)
	UpdatePost(uint32, uint32, string) (int64, error)
	DeletePost(uint32) error
}

type postUsecase struct {
	postRepository  repository.PostRepository
	imageRepository repository.ImageRepository
}

// NewPostUsecase Postデータに関するUsecaseを生成
func NewPostUsecase(pr repository.PostRepository, ir repository.ImageRepository) PostUsecase {
	return &postUsecase{
		postRepository:  pr,
		imageRepository: ir,
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

func (pu *postUsecase) GetPosts(sid uint32) ([]postsResponse, error) {
	var responses []postsResponse

	posts, err := pu.postRepository.FindAll(sid)
	if err != nil {
		return []postsResponse{}, err
	}
	// 投稿が存在する場合
	if len(*posts) > 0 {
		for i := 0; i < len(*posts); i++ {
			// ユーザー値取得
			user, err := pu.postRepository.FindByUserID((*posts)[i].UserID)
			if err != nil {
				return []postsResponse{}, err
			}
			// ユーザー画像取得
			img, err := pu.imageRepository.FindByID((*posts)[i].UserID)
			if err != nil {
				return []postsResponse{}, err
			}
			err = pu.imageRepository.DownloadS3(img)
			if err != nil {
				return []postsResponse{}, err
			}
			uri, err := security.Base64EncodeToString(img.Buf)
			if err != nil {
				return []postsResponse{}, err
			}
			// コメント数取得
			commentsCount := pu.postRepository.FindCommentsCount((*posts)[i].ID)

			if err != nil {
				return []postsResponse{}, err
			}

			res := postsResponse{
				ID:            (*posts)[i].ID,
				Text:          (*posts)[i].Text,
				Rating:        (*posts)[i].Rating,
				UserID:        (*posts)[i].UserID,
				UserNickname:  (*user).Nickname,
				UserImage:     uri,
				CommentsCount: commentsCount,
			}
			responses = append(responses, res)
		}
	}
	return responses, nil
}

func (pu *postUsecase) GetPost(sid, pid uint32) (*models.Post, error) {
	post, err := pu.postRepository.FindByID(sid, pid)
	if err != nil {
		return &models.Post{}, err
	}
	return post, nil
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

type postsResponse struct {
	ID            uint32 `json:"id"`
	Text          string `json:"text"`
	Rating        uint32 `json:"rating"`
	UserID        uint32 `json:"user_id"`
	UserNickname  string `json:"user_nickname"`
	UserImage     string `json:"user_image"`
	CommentsCount uint32 `json:"comments_count"`
}
