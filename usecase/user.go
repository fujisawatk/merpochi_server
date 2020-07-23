package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/usecase/validations"
	"merpochi_server/util/security"
)

// UserUsecase Userに対するUsecaseのインターフェイス
type UserUsecase interface {
	GetUsers() ([]models.User, error)
	CreateUser(nickname, email, password string) (models.User, error)
	GetUser(uint32) (models.User, error)
	UpdateUser(uint32, string, string) (int64, error)
	DeleteUser(uint32) error
	CommentedShops(uint32) ([]models.Shop, error)
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
	user := models.User{
		Nickname: nickname,
		Email:    email,
		Password: password,
	}

	err := validations.UserCreateValidate(&user)
	if err != nil {
		return models.User{}, err
	}

	// メールアドレスの重複確認
	err = uu.userRepository.SearchUser(user.Email)
	if err != nil {
		return models.User{}, err
	}

	// パスワードのハッシュ化
	var hashedPassword []byte
	hashedPassword, err = security.Hash(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)

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

	err := validations.UserUpdateValidate(&user)
	if err != nil {
		return 0, err
	}

	rows, err := uu.userRepository.Update(uid, user)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (uu userUsecase) DeleteUser(uid uint32) error {
	_, err := uu.userRepository.Delete(uid)
	if err != nil {
		return err
	}
	return nil
}

func (uu userUsecase) CommentedShops(uid uint32) ([]models.Shop, error) {
	shops, err := uu.userRepository.FindShops(uid)
	if err != nil {
		return []models.Shop{}, err
	}
	return shops, nil
}
