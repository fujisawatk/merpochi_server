package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// FavoriteUsecase Favoriteに対するUsecaseのインターフェイス
type FavoriteUsecase interface {
	CreateFavorite(uint32, uint32) (models.Favorite, error)
	DeleteFavorite(uint32, uint32) error
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

func (fu favoriteUsecase) CreateFavorite(sid uint32, uid uint32) (models.Favorite, error) {
	favorite := models.Favorite{
		UserID: uid,
		ShopID: sid,
	}

	favorite, err := fu.favoriteRepository.Save(favorite)
	if err != nil {
		return models.Favorite{}, err
	}

	// お気に入りしたユーザー値を取得
	favoriteUser, err := fu.favoriteRepository.FindFavoriteUser(favorite.UserID)
	if err != nil {
		return models.Favorite{}, err
	}
	favorite.User = favoriteUser

	return favorite, nil
}

func (fu favoriteUsecase) DeleteFavorite(sid uint32, uid uint32) error {
	_, err := fu.favoriteRepository.Delete(sid, uid)
	if err != nil {
		return err
	}
	return nil
}
