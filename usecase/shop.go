package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// ShopUsecase Shopに対するUsecaseのインターフェイス
type ShopUsecase interface {
	GetShops([]string) ([]shopResponse, error)
	CreateShop(models.Shop) (models.Shop, error)
}

type shopUsecase struct {
	shopRepository repository.ShopRepository
}

// NewShopUsecase Shopデータに関するUsecaseを生成
func NewShopUsecase(sr repository.ShopRepository) ShopUsecase {
	return &shopUsecase{
		shopRepository: sr,
	}
}

func (su shopUsecase) GetShops(shopCodes []string) ([]shopResponse, error) {
	var counts []shopResponse

	// 取得した店舗IDを1件ずつ登録されているか確認
	for _, code := range shopCodes {
		var res shopResponse
		shop, err := su.shopRepository.SearchShop(code)
		// 登録されていない場合
		if err != nil {
			res = shopResponse{
				ID:             0,
				CommentsCount:  0,
				FavoritesCount: 0,
			}
			counts = append(counts, res)
		} else {
			// 登録されている場合
			commentsCount, err := su.shopRepository.FindCommentsCount(shop.ID)
			// 登録後にコメントが削除された場合
			if err != nil {
				commentsCount = 0
			}
			favoritesCount, err := su.shopRepository.FindFavoritesCount(shop.ID)
			// 登録後にいいねが削除された場合
			if err != nil {
				favoritesCount = 0
			}
			res = shopResponse{
				ID:             shop.ID,
				CommentsCount:  int(commentsCount),
				FavoritesCount: int(favoritesCount),
			}
			counts = append(counts, res)
		}
	}
	return counts, nil
}

func (su shopUsecase) CreateShop(req models.Shop) (models.Shop, error) {
	shop, err := su.shopRepository.Save(req)
	if err != nil {
		return models.Shop{}, err
	}
	return shop, nil
}

type shopResponse struct {
	ID             uint32 `json:"id"`
	CommentsCount  int    `json:"comments_count"`
	FavoritesCount int    `json:"favorites_count"`
}
