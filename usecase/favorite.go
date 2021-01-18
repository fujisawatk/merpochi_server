package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// FavoriteUsecase Favoriteに対するUsecaseのインターフェイス
type FavoriteUsecase interface {
	GetFavorites(uint32) ([]models.Favorite, error)
	CreateFavorite(uint32, uint32) (models.Favorite, error)
	DeleteFavorite(uint32, uint32) error
}

type favoriteUsecase struct {
	favoriteRepository repository.FavoriteRepository
	userRepository     repository.UserRepository
}

// NewFavoriteUsecase Favoriteデータに関するUsecaseを生成
func NewFavoriteUsecase(
	fr repository.FavoriteRepository,
	ur repository.UserRepository,
) FavoriteUsecase {
	return &favoriteUsecase{
		favoriteRepository: fr,
		userRepository:     ur,
	}
}

func (fu favoriteUsecase) GetFavorites(sid uint32) ([]models.Favorite, error) {
	favorites, err := fu.favoriteRepository.FindAll(sid)
	if err != nil {
		return []models.Favorite{}, err
	}
	// お気に入りが存在する場合
	if len(favorites) > 0 {
		// 取得した店舗のお気に入りに紐付くユーザーを取得
		for i := 0; i < len(favorites); i++ {
			favoriteUser, err := fu.userRepository.FindByID(favorites[i].UserID)
			if err != nil {
				return []models.Favorite{}, err
			}
			favorites[i].User = favoriteUser
		}
	}
	return favorites, nil
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
	favoriteUser, err := fu.userRepository.FindByID(favorite.UserID)
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
