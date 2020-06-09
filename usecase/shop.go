package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// ShopUsecase Shopに対するUsecaseのインターフェイス
type ShopUsecase interface {
	GetShops() ([]models.Shop, error)
	CreateShop(string) (models.Shop, error)
	GetShop(uint32) (models.Shop, error)
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

func (su shopUsecase) GetShops() ([]models.Shop, error) {
	shops, err := su.shopRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return shops, nil
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

func (su shopUsecase) GetShop(uid uint32) (models.Shop, error) {
	shop, err := su.shopRepository.FindByID(uid)
	if err != nil {
		return models.Shop{}, err
	}
	return shop, nil
}
