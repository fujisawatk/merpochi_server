package usecase

import (
	"errors"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/interfaces/validations"
)

// UserUsecase Userに対するUsecaseのインターフェイス
type UserUsecase interface {
	GetUsers() ([]models.User, error)
	CreateUser(nickname, email, password string) (models.User, error)
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

func (uu userUsecase) CreateUser(nickname, email, password string) (models.User, error) {
	var err error

	user := models.User{
		Nickname: nickname,
		Email:    email,
		Password: password,
	}

	errs := validations.UserCreateValidate(&user)
	if errs != nil {
		return models.User{}, errors.New("validation error")
	}

	user, err = uu.userRepository.Save(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
