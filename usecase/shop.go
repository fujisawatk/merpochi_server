package usecase

import (
	"fmt"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// ShopUsecase Shopに対するUsecaseのインターフェイス
type ShopUsecase interface {
	GetShops([]string) ([]shopResponse, error)
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

func (su shopUsecase) GetShops(shopCodes []string) ([]shopResponse, error) {
	var commentsCount []shopResponse

	// 取得した店舗IDを1件ずつ登録されているか確認
	for _, code := range shopCodes {
		var res shopResponse
		shop, err := su.shopRepository.SearchShop(code)
		// 登録されていない場合
		if err != nil {
			res = shopResponse{
				ID:    0,
				Count: 0,
			}
			commentsCount = append(commentsCount, res)
		} else {
			// 登録されている場合
			count, err := su.shopRepository.FindAll(shop.ID)
			// 登録後にコメントが削除された場合
			if err != nil {
				count = 0
			}
			res = shopResponse{
				ID:    shop.ID,
				Count: int(count),
			}
			commentsCount = append(commentsCount, res)
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
	// コメントが存在する場合
	if len(comment) > 0 {
		// 取得した店舗のコメントに紐付くユーザーを取得
		for i := 0; i < len(comment); i++ {
			fmt.Println(comment[i])
			commentUser, err := su.shopRepository.FindCommentUser(comment[i].UserID)
			if err != nil {
				return []models.Comment{}, err
			}
			comment[i].User = commentUser
		}
	}
	return comment, nil
}

type shopResponse struct {
	ID    uint32 `json:"id"`
	Count int    `json:"count"`
}
