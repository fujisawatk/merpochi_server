package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// ShopUsecase Shopに対するUsecaseのインターフェイス
type ShopUsecase interface {
	SearchShops([]string, uint32) ([]searchShopsResponse, error)
	CreateShop(models.Shop) (models.Shop, error)
	GetShop(string) (*models.Shop, error)
	GetPostedShop(uint32) (*models.Shop, error)
}

type shopUsecase struct {
	shopRepository     repository.ShopRepository
	postRepository     repository.PostRepository     // has many
	favoriteRepository repository.FavoriteRepository // has many
	bookmarkRepository repository.BookmarkRepository // has many
}

// NewShopUsecase Shopデータに関するUsecaseを生成
func NewShopUsecase(
	sr repository.ShopRepository,
	pr repository.PostRepository,
	fr repository.FavoriteRepository,
	br repository.BookmarkRepository,
) ShopUsecase {
	return &shopUsecase{
		shopRepository:     sr,
		postRepository:     pr,
		favoriteRepository: fr,
		bookmarkRepository: br,
	}
}

func (su *shopUsecase) SearchShops(shopCodes []string, uid uint32) ([]searchShopsResponse, error) {
	var counts []searchShopsResponse

	// 取得した店舗IDを1件ずつ登録されているか確認
	for _, code := range shopCodes {
		var res searchShopsResponse
		shop, err := su.shopRepository.FindByCode(code)
		// 登録されていない場合
		if err != nil {
			res = searchShopsResponse{
				ID:             0,
				RatingCount:    0,
				BookmarksCount: 0,
				BookmarkUser:   false,
			}
			counts = append(counts, res)
		} else {
			// 評価が4以上である投稿数を取得
			postsCount := su.postRepository.CountByShopID(shop.ID)
			// ブックマーク数を取得
			bookmarksCount := su.bookmarkRepository.CountByShopID(shop.ID)
			// お気に入り（リピートしたいボタン）が押された数を取得
			favoritesCount := su.favoriteRepository.CountByShopID(shop.ID)
			// APIを呼び出したユーザーがブックマークしているか確認
			bookmarkUser := su.bookmarkRepository.SearchUser(shop.ID, uid)
			// APIを呼び出したユーザーがお気に入りしているか確認
			favoriteUser := su.favoriteRepository.SearchUser(shop.ID, uid)
			res = searchShopsResponse{
				ID:             shop.ID,
				RatingCount:    int(postsCount) + int(favoritesCount),
				BookmarksCount: int(bookmarksCount),
				BookmarkUser:   bookmarkUser,
				FavoritesCount: int(favoritesCount),
				FavoriteUser:   favoriteUser,
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

func (su *shopUsecase) GetShop(code string) (*models.Shop, error) {
	shop, err := su.shopRepository.FindByCode(code)
	if err != nil {
		return &models.Shop{}, err
	}
	return shop, nil
}

func (su *shopUsecase) GetPostedShop(pid uint32) (*models.Shop, error) {
	shop, err := su.shopRepository.FindByPostID(pid)
	if err != nil {
		return &models.Shop{}, err
	}
	return shop, nil
}

type searchShopsResponse struct {
	ID             uint32 `json:"id"`
	RatingCount    int    `json:"rating_count"`
	BookmarksCount int    `json:"bookmarks_count"`
	BookmarkUser   bool   `json:"bookmark_user"`
	FavoritesCount int    `json:"favorites_count"`
	FavoriteUser   bool   `json:"favorite_user"`
}
