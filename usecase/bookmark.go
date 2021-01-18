package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// BookmarkUsecase Bookmarkに対するUsecaseのインターフェイス
type BookmarkUsecase interface {
	CreateBookmark(uint32, uint32) (*models.Bookmark, error)
	DeleteBookmark(uint32, uint32) error
}

type bookmarkUsecase struct {
	bookmarkRepository repository.BookmarkRepository
}

// NewBookmarkUsecase Bookmarkデータに関するUsecaseを生成
func NewBookmarkUsecase(br repository.BookmarkRepository) BookmarkUsecase {
	return &bookmarkUsecase{
		bookmarkRepository: br,
	}
}

func (bu *bookmarkUsecase) CreateBookmark(sid uint32, uid uint32) (*models.Bookmark, error) {
	bookmark := &models.Bookmark{
		UserID: uid,
		ShopID: sid,
	}

	bookmark, err := bu.bookmarkRepository.Save(bookmark)
	if err != nil {
		return &models.Bookmark{}, err
	}

	return bookmark, nil
}

func (bu bookmarkUsecase) DeleteBookmark(sid uint32, uid uint32) error {
	err := bu.bookmarkRepository.Delete(sid, uid)
	if err != nil {
		return err
	}
	return nil
}
