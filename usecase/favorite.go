package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// FavoriteUsecase Favoriteに対するUsecaseのインターフェイス
type FavoriteUsecase interface {
	CreateFavorite(uint32, uint32) (uint32, error)
	// DeleteFavorite(uint32) error
}

type favoriteUsecase struct {
	favoriteRepository repository.FavoriteRepository
}

// NewFavoriteUsecase Favoriteデータに関するUsecaseを生成
func NewFavoriteUsecase(fr repository.FavoriteRepository) FavoriteUsecase {
	return &favoriteUsecase{
		favoriteRepository: fr,
	}
}

func (fu favoriteUsecase) CreateFavorite(sid uint32, uid uint32) (uint32, error) {
	favorite := models.Favorite{
		UserID: uid,
		ShopID: sid,
	}

	err := fu.favoriteRepository.Save(favorite)
	if err != nil {
		return 0, err
	}

	// お気に入り登録したページのお気に入り総数を取得
	count, err := fu.favoriteRepository.Search(favorite.ShopID)
	if err != nil {
		return 0, err
	}
	return count, nil
}
