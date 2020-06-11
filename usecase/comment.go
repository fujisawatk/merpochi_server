package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// CommentUsecase Commentに対するUsecaseのインターフェイス
type CommentUsecase interface {
	CreateComment(string, uint32) (models.Comment, error)
	UpdateComment(uint32, string) (int64, error)
	DeleteComment(uint32) error
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

func (cu commentUsecase) UpdateComment(cid uint32, text string) (int64, error) {
	comment := models.Comment{
		Text: text,
	}

	rows, err := cu.commentRepository.Update(cid, comment)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (cu commentUsecase) DeleteComment(cid uint32) error {
	_, err := cu.commentRepository.Delete(cid)
	if err != nil {
		return err
	}
	return nil
}
