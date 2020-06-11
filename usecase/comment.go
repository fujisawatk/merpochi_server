package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// CommentUsecase Commentに対するUsecaseのインターフェイス
type CommentUsecase interface {
	CreateComment(string, uint32) (models.Comment, error)
}

type commentUsecase struct {
	commentRepository repository.CommentRepository
}

// NewCommentUsecase Commentデータに関するUsecaseを生成
func NewCommentUsecase(cr repository.CommentRepository) CommentUsecase {
	return &commentUsecase{
		commentRepository: cr,
	}
}

func (cu commentUsecase) CreateComment(text string, sid uint32) (models.Comment, error) {
	var err error
	comment := models.Comment{
		Text:   text,
		ShopID: sid,
	}

	comment, err = cu.commentRepository.Save(comment)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}
