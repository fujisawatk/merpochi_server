package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// ShopUsecase Shopに対するUsecaseのインターフェイス
type ShopUsecase interface {
	SearchShops([]string) ([]searchShopsResponse, error)
	CreateShop(models.Shop) (models.Shop, error)
	MeShops(uint32) (meShopsResponse, error)
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

func (su shopUsecase) SearchShops(shopCodes []string) ([]searchShopsResponse, error) {
	var counts []searchShopsResponse

	// 取得した店舗IDを1件ずつ登録されているか確認
	for _, code := range shopCodes {
		var res searchShopsResponse
		shop, err := su.shopRepository.Search(code)
		// 登録されていない場合
		if err != nil {
			res = searchShopsResponse{
				ID:             0,
				CommentsCount:  0,
				FavoritesCount: 0,
			}
			counts = append(counts, res)
		} else {
			// 登録されている場合
			commentsCount := su.shopRepository.FindCommentsCount(shop.ID)
			// 登録後にコメントが削除された場合
			if err != nil {
				commentsCount = 0
			}
			favoritesCount := su.shopRepository.FindFavoritesCount(shop.ID)
			// 登録後にいいねが削除された場合
			if err != nil {
				favoritesCount = 0
			}
			res = searchShopsResponse{
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

func (su shopUsecase) MeShops(uid uint32) (meShopsResponse, error) {
	commentedShops, err := su.shopRepository.FindCommentedShops(uid)
	if err != nil {
		return meShopsResponse{}, err
	}

	commentedShops = DelDuplicateShops(commentedShops)
	if err != nil {
		return meShopsResponse{}, err
	}

	favoritedShops, err := su.shopRepository.FindFavoritedShops(uid)
	if err != nil {
		return meShopsResponse{}, err
	}

	favoritedShops = DelDuplicateShops(favoritedShops)
	if err != nil {
		return meShopsResponse{}, err
	}

	res := meShopsResponse{
		CommentedShops: commentedShops,
		FavoritedShops: favoritedShops,
	}

	return res, nil
}

// DelDuplicateShops 店舗情報重複削除
func DelDuplicateShops(shops []models.Shop) []models.Shop {
	m := make(map[string]bool)
	uniq := []models.Shop{}

	for _, shop := range shops {
		if !m[shop.Name] {
			m[shop.Name] = true
			uniq = append(uniq, shop)
		}
	}
	return uniq
}

type searchShopsResponse struct {
	ID             uint32 `json:"id"`
	CommentsCount  int    `json:"comments_count"`
	FavoritesCount int    `json:"favorites_count"`
}

type meShopsResponse struct {
	CommentedShops []models.Shop `json:"commented_shops"`
	FavoritedShops []models.Shop `json:"favorited_shops"`
}
