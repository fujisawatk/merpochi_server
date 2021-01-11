package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// ShopUsecase Shopに対するUsecaseのインターフェイス
type ShopUsecase interface {
	SearchShops([]string) ([]searchShopsResponse, error)
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

func (su *shopUsecase) SearchShops(shopCodes []string) ([]searchShopsResponse, error) {
	var counts []searchShopsResponse

	// 取得した店舗IDを1件ずつ登録されているか確認
	for _, code := range shopCodes {
		var res searchShopsResponse
		shop, err := su.shopRepository.Search(code)
		// 登録されていない場合
		if err != nil {
			res = searchShopsResponse{
				ID:             0,
				RatingCount:    0,
				BookmarksCount: 0,
			}
			counts = append(counts, res)
		} else {
			// 評価が4以上である投稿数を取得
			postsCount := su.shopRepository.FindPostsCount(shop.ID)
			// お気に入り（リピートしたいボタン）が押された数を取得
			favoritesCount := su.shopRepository.FindFavoritesCount(shop.ID)
			// ブックマーク数を取得
			bookmarksCount := su.shopRepository.FindBookmarksCount(shop.ID)
			res = searchShopsResponse{
				ID:             shop.ID,
				RatingCount:    int(postsCount) + int(favoritesCount),
				BookmarksCount: int(bookmarksCount),
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

type searchShopsResponse struct {
	ID             uint32 `json:"id"`
	RatingCount    int    `json:"rating_count"`
	BookmarksCount int    `json:"bookmarks_count"`
}
