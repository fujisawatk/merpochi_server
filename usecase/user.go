package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/usecase/validations"
	"merpochi_server/util/security"
)

// UserUsecase Userに対するUsecaseのインターフェイス
type UserUsecase interface {
	CreateUser(string, string, string, string) (models.User, error)
	GetUser(uint32) (models.User, error)
	UpdateUser(uint32, string, string, string) (int64, error)
	DeleteUser(uint32) error
	MylistUser(uint32) (*userResponse, error)
	MeUser(uint32) (*meUserResponse, error)
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

func (uu userUsecase) CreateUser(nickname, email, password, genre string) (models.User, error) {
	user := models.User{
		Nickname: nickname,
		Email:    email,
		Password: password,
		Genre:    genre,
	}

	err := validations.UserCreateValidate(&user)
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

func (uu userUsecase) UpdateUser(uid uint32, nickname, email, genre string) (int64, error) {
	user := models.User{
		Nickname: nickname,
		Email:    email,
		Genre:    genre,
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

func (uu *userUsecase) MylistUser(uid uint32) (*userResponse, error) {
	bookmarkedShops, err := uu.userRepository.FindBookmarkedShops(uid)
	if err != nil {
		return &userResponse{}, err
	}

	favoritedShops, err := uu.userRepository.FindFavoritedShops(uid)
	if err != nil {
		return &userResponse{}, err
	}

	res := &userResponse{
		BookmarkedShops: *(bookmarkedShops),
		FavoritedShops:  *(favoritedShops),
	}
	return res, nil
}

func (uu *userUsecase) MeUser(uid uint32) (*meUserResponse, error) {
	myPosts, err := uu.userRepository.FindMyPosts(uid)
	if err != nil {
		return &meUserResponse{}, err
	}

	commentedPosts, err := uu.userRepository.FindCommentedPosts(uid)
	if err != nil {
		return &meUserResponse{}, err
	}
	uniqPosts := DelDuplicatePosts(commentedPosts)

	res := &meUserResponse{
		MyPosts:        *(myPosts),
		CommentedPosts: uniqPosts,
	}
	return res, nil
}

// DelDuplicatePosts 店舗情報重複削除
func DelDuplicatePosts(posts *[]models.Post) []models.Post {
	m := make(map[uint32]bool)
	uniq := []models.Post{}

	for _, post := range *(posts) {
		if !m[post.ShopID] {
			m[post.ShopID] = true
			uniq = append(uniq, post)
		}
	}
	return uniq
}

type userResponse struct {
	BookmarkedShops []models.Shop `json:"bookmarked_shops"`
	FavoritedShops  []models.Shop `json:"favorited_shops"`
}

type meUserResponse struct {
	MyPosts        []models.Post `json:"my_posts"`
	CommentedPosts []models.Post `json:"commented_posts"`
}
