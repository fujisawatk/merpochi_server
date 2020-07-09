package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// ShopUsecase Shopに対するUsecaseのインターフェイス
type ShopUsecase interface {
	GetShops([]string) ([]int, error)
	CreateShop(string) (models.Shop, error)
	GetShop(uint32) ([]models.Comment, error)
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

func (su shopUsecase) GetShops(shopCodes []string) ([]int, error) {
	var commentsCount []int

	// 取得した店舗IDを1件ずつ登録されているか確認
	for _, code := range shopCodes {
		shop, err := su.shopRepository.SearchShop(code)
		// 登録されていない場合
		if err != nil {
			commentsCount = append(commentsCount, 0)
		} else {
			// 登録されている場合
			count, err := su.shopRepository.FindAll(shop.ID)
			// 登録後にコメントが削除された場合
			if err != nil {
				count = 0
			}
			commentsCount = append(commentsCount, int(count))
		}
	}
	return commentsCount, nil
}

func (su shopUsecase) CreateShop(code string) (models.Shop, error) {
	var err error

	shop := models.Shop{
		Code: code,
	}

	shop, err = su.shopRepository.Save(shop)
	if err != nil {
		return models.Shop{}, err
	}
	return shop, nil
}

func (su shopUsecase) GetShop(sid uint32) ([]models.Comment, error) {
	comment, err := su.shopRepository.FindByID(sid)
	if err != nil {
		return []models.Comment{}, err
	}
	return comment, nil
}
