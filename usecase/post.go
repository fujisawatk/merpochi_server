package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/usecase/validations"
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
	postRepository repository.PostRepository
}

// NewPostUsecase Postデータに関するUsecaseを生成
func NewPostUsecase(pr repository.PostRepository) PostUsecase {
	return &postUsecase{
		postRepository: pr,
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
		// 投稿のコメント数を取得
		for i := 0; i < len(*posts); i++ {
			user, err := pu.postRepository.FindByUserID((*posts)[i].UserID)
			if err != nil {
				return []postsResponse{}, err
			}
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
	CommentsCount uint32 `json:"comments_count"`
}
