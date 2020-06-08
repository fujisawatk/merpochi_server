package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// ShopUsecase Shopに対するUsecaseのインターフェイス
type ShopUsecase interface {
	GetShops() ([]models.Shop, error)
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
