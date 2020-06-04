package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// UserUsecase Userに対するUsecaseのインターフェイス
type UserUsecase interface {
	GetUsers() ([]models.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

// NewUserUsecase Userデータに関するUsecaseを生成
func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu userUsecase) GetUsers() ([]models.User, error) {
	users, err := uu.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
