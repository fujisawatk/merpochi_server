package usecase

import (
	"fmt"
	"merpochi_server/config"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/usecase/validations"

	"github.com/dgrijalva/jwt-go"
)

// AuthUsecase ユーザー認証に対するUsecaseのインターフェイス
type AuthUsecase interface {
	LoginUser(string, string) (string, error)
	VerifyUser(string) (authResponse, error)
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

	err := validations.UserLoginValidate(&user)
	if err != nil {
		return "", err
	}

	token, err := au.authRepository.SignIn(user.Email, user.Password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (au authUsecase) VerifyUser(authToken string) (authResponse, error) {
	token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("トークン系エラーです")
		}
		return config.SECRETKEY, nil
	})
	if error != nil {
		fmt.Println(error)
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	uid := claims["user_id"].(float64)

	user, rToken, err := au.authRepository.FindCurrentUser(uint32(uid))
	if err != nil {
		return authResponse{}, err
	}
	cUser := authResponse{
		ID:       user.ID,
		Nickname: user.Nickname,
		Email:    user.Email,
		Token:    rToken,
	}
	return cUser, nil
}

type authResponse struct {
	ID       uint32 `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
