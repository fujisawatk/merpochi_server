package usecase

import (
	"bytes"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/usecase/validations"
	"merpochi_server/util/security"
	"time"
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
	userRepository    repository.UserRepository
	shopRepository    repository.ShopRepository
	postRepository    repository.PostRepository    // has many
	commentRepository repository.CommentRepository // has many
	imageRepository   repository.ImageRepository   // has many
}

// NewUserUsecase Userデータに関するUsecaseを生成
func NewUserUsecase(
	ur repository.UserRepository,
	sr repository.ShopRepository,
	pr repository.PostRepository,
	cr repository.CommentRepository,
	ir repository.ImageRepository,
) UserUsecase {
	return &userUsecase{
		userRepository:    ur,
		shopRepository:    sr,
		postRepository:    pr,
		commentRepository: cr,
		imageRepository:   ir,
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
	bookmarkedShops, err := uu.shopRepository.FindBookmarkedShops(uid)
	if err != nil {
		return &userResponse{}, err
	}

	favoritedShops, err := uu.shopRepository.FindFavoritedShops(uid)
	if err != nil {
		return &userResponse{}, err
	}

	res := &userResponse{
		BookmarkedShops: *(bookmarkedShops),
		FavoritedShops:  *(favoritedShops),
	}
	return res, nil
}

// ※関数分けたほうがいい
func (uu *userUsecase) MeUser(uid uint32) (*meUserResponse, error) {
	// ログインユーザーが投稿したレビュー
	myPosts, err := uu.postRepository.FindByUserID(uid)
	if err != nil {
		return &meUserResponse{}, err
	}
	var myPostsData []postData
	if len(*myPosts) > 0 {
		// 投稿したユーザー情報を取得
		for i := 0; i < len(*myPosts); i++ {
			postedUser, imgURI, time, err := uu.GetUserData((*myPosts)[i].UserID, (*myPosts)[i].CreatedAt, (*myPosts)[i].UpdatedAt)
			if err != nil {
				return &meUserResponse{}, err
			}
			// コメント数取得
			commentsCount := uu.commentRepository.CountByPostID((*myPosts)[i].ID)

			imgs, err := uu.GetPostImage((*myPosts)[i].UserID, (*myPosts)[i].ShopID, (*myPosts)[i].ID)
			if err != nil {
				return &meUserResponse{}, err
			}

			res := postData{
				ID:            (*myPosts)[i].ID,
				Text:          (*myPosts)[i].Text,
				Rating:        (*myPosts)[i].Rating,
				Images:        imgs,
				UserID:        (*myPosts)[i].UserID,
				UserNickname:  postedUser,
				UserImage:     imgURI,
				CommentsCount: commentsCount,
				Time:          time,
			}
			myPostsData = append(myPostsData, res)
		}
	}
	// ログインユーザーがコメントしたレビュー
	commentedPosts, err := uu.postRepository.FindCommentedPosts(uid)
	if err != nil {
		return &meUserResponse{}, err
	}
	var commentedPostsData []postData
	if len(*commentedPosts) > 0 {
		// 投稿したユーザー情報を取得
		for i := 0; i < len(*commentedPosts); i++ {
			postedUser, imgURI, time, err := uu.GetUserData((*commentedPosts)[i].UserID, (*commentedPosts)[i].CreatedAt, (*commentedPosts)[i].UpdatedAt)
			if err != nil {
				return &meUserResponse{}, err
			}
			// コメント数取得
			commentsCount := uu.commentRepository.CountByPostID((*commentedPosts)[i].ID)

			imgs, err := uu.GetPostImage((*commentedPosts)[i].UserID, (*commentedPosts)[i].ShopID, (*commentedPosts)[i].ID)
			if err != nil {
				return &meUserResponse{}, err
			}

			res := postData{
				ID:            (*commentedPosts)[i].ID,
				Text:          (*commentedPosts)[i].Text,
				Rating:        (*commentedPosts)[i].Rating,
				Images:        imgs,
				UserID:        (*commentedPosts)[i].UserID,
				UserNickname:  postedUser,
				UserImage:     imgURI,
				CommentsCount: commentsCount,
				Time:          time,
			}
			commentedPostsData = append(commentedPostsData, res)
		}
	}
	uniqPosts := uu.DelDuplicatePosts(commentedPostsData)

	res := &meUserResponse{
		MyPosts:        myPostsData,
		CommentedPosts: uniqPosts,
	}
	return res, nil
}

// ユーザー情報取得〜整形まで
func (uu *userUsecase) GetUserData(uid uint32, createdAt, updatedAt time.Time) (string, string, string, error) {
	user, err := uu.userRepository.FindByID(uid)
	if err != nil {
		return "", "", "", err
	}

	imgURI, err := uu.GetUserImage(uid)
	if err != nil {
		return "", "", "", err
	}
	// 作成or編集時刻整形
	var time string
	format1 := "2006/01/02 15:04:05"
	if createdAt != updatedAt {
		time = "編集済 " + (updatedAt).Format(format1)
	} else {
		time = (createdAt).Format(format1)
	}
	return user.Nickname, imgURI, time, nil
}

// ユーザー画像取得〜base64エンコード文字列生成まで
func (uu *userUsecase) GetUserImage(uid uint32) (string, error) {
	img, err := uu.imageRepository.FindByUserID(uid)
	if err != nil {
		return "", err
	}

	err = uu.imageRepository.DownloadS3(img, "merpochi-users-image")
	if err != nil {
		return "", err
	}

	uri, err := security.Base64EncodeToString(img.Buf)
	if err != nil {
		return "", err
	}
	return uri, nil
}

// 投稿画像取得〜base64エンコード文字列生成まで
func (uu *userUsecase) GetPostImage(uid, sid, pid uint32) ([]imageData, error) {
	imgs, err := uu.imageRepository.FindAll(uid, sid, pid)
	if err != nil {
		return []imageData{}, err
	}
	var responses []imageData
	if len(*imgs) > 0 {
		for i := 0; i < len(*imgs); i++ {
			img := &models.Image{
				Name: (*imgs)[i].Name,
				Buf:  &bytes.Buffer{},
			}
			err = uu.imageRepository.DownloadS3(img, "merpochi-posts-image")
			if err != nil {
				return []imageData{}, err
			}

			uri, err := security.Base64EncodeToString((*img).Buf)
			if err != nil {
				return []imageData{}, err
			}
			res := imageData{
				ID:  (*imgs)[i].ID,
				URI: uri,
			}
			responses = append(responses, res)
		}
	}
	return responses, nil
}

// DelDuplicatePosts 投稿情報重複削除
func (uu *userUsecase) DelDuplicatePosts(posts []postData) []postData {
	m := make(map[uint32]bool)
	uniq := []postData{}

	for _, post := range posts {
		if !m[post.ID] {
			m[post.ID] = true
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
	MyPosts        []postData `json:"my_posts"`
	CommentedPosts []postData `json:"commented_posts"`
}

type postData struct {
	ID            uint32      `json:"id"`
	Text          string      `json:"text"`
	Rating        uint32      `json:"rating"`
	Images        []imageData `json:"images"`
	UserID        uint32      `json:"user_id"`
	UserNickname  string      `json:"user_nickname"`
	UserImage     string      `json:"user_image"`
	CommentsCount uint32      `json:"comments_count"`
	Time          string      `json:"time"`
}
