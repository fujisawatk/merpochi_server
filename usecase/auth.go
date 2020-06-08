package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// AuthUsecase ユーザー認証に対するUsecaseのインターフェイス
type AuthUsecase interface {
	LoginUser(string, string) (string, error)
}

type authUsecase struct {
	authRepository repository.AuthRepository
}

// NewAuthUsecase ユーザー認証に関するUsecaseを生成
func NewAuthUsecase(ar repository.AuthRepository) AuthUsecase {
	return &authUsecase{
		authRepository: ar,
	}
}

func (au authUsecase) LoginUser(email, password string) (string, error) {
	user := models.User{
		Email:    email,
		Password: password,
	}

	token, err := au.authRepository.SignIn(user.Email, user.Password)
	if err != nil {
		return "", err
	}
	return token, nil
}
