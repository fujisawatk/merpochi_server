package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/usecase/validations"
)

// PostUsecase Postに対するUsecaseのインターフェイス
type PostUsecase interface {
	CreatePost(string, uint32, uint32, uint32) (*models.Post, error)
	GetPosts(uint32) (*[]models.Post, error)
	GetPost(uint32, uint32) (*models.Post, error)
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

	err := validations.PostCreateValidate(post)
	if err != nil {
		return &models.Post{}, err
	}

	err = pu.postRepository.Save(post)
	if err != nil {
		return &models.Post{}, err
	}
	return post, nil
}

func (pu *postUsecase) GetPosts(sid uint32) (*[]models.Post, error) {
	posts, err := pu.postRepository.FindAll(sid)
	if err != nil {
		return &[]models.Post{}, err
	}
	return posts, nil
}

func (pu *postUsecase) GetPost(sid, pid uint32) (*models.Post, error) {
	post, err := pu.postRepository.FindByID(sid, pid)
	if err != nil {
		return &models.Post{}, err
	}
	return post, nil
}
