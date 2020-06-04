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
	GetUser(uint32) (models.User, error)
	UpdateUser(uint32, string, string) (int64, error)
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

func (uu userUsecase) GetUser(uid uint32) (models.User, error) {
	user, err := uu.userRepository.FindByID(uid)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (uu userUsecase) UpdateUser(uid uint32, nickname, email string) (int64, error) {
	user := models.User{
		Nickname: nickname,
		Email:    email,
	}

	// errs := validations.UserUpdateValidate(&user)
	// if errs != nil {
	// 	return 0, errors.New("validation error")
	// }

	rows, err := uu.userRepository.Update(uid, user)
	if err != nil {
		return 0, err
	}
	return rows, nil
}
