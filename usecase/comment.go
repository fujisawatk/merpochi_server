package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/usecase/validations"
	"merpochi_server/util/security"
	"time"
)

// CommentUsecase Commentに対するUsecaseのインターフェイス
type CommentUsecase interface {
	GetComments(uint32) ([]commentResponse, error)
	CreateComment(string, uint32, uint32) (commentResponse, error)
	UpdateComment(uint32, string) (int64, error)
	DeleteComment(uint32) error
}

type commentUsecase struct {
	commentRepository repository.CommentRepository
	userRepository    repository.UserRepository // belongs to
	postRepository    repository.PostRepository // belongs to
	imageRepository   repository.ImageRepository
}

// NewCommentUsecase Commentデータに関するUsecaseを生成
func NewCommentUsecase(
	cr repository.CommentRepository,
	ur repository.UserRepository,
	pr repository.PostRepository,
	ir repository.ImageRepository,
) CommentUsecase {
	return &commentUsecase{
		commentRepository: cr,
		userRepository:    ur,
		postRepository:    pr,
		imageRepository:   ir,
	}
}

func (cu *commentUsecase) GetComments(pid uint32) ([]commentResponse, error) {
	comments, err := cu.commentRepository.FindAll(pid)
	if err != nil {
		return []commentResponse{}, err
	}

	var responses []commentResponse
	// コメントが存在する場合
	if len(*comments) > 0 {
		// 取得した投稿にコメントしたユーザー情報を取得
		for i := 0; i < len(*comments); i++ {
			commentedUser, imgURI, time, err := cu.GetUserData((*comments)[i].UserID, (*comments)[i].CreatedAt, (*comments)[i].UpdatedAt)
			if err != nil {
				return []commentResponse{}, err
			}
			res := commentResponse{
				ID:           (*comments)[i].ID,
				Text:         (*comments)[i].Text,
				UserID:       (*comments)[i].UserID,
				UserNickname: commentedUser,
				UserImage:    imgURI,
				Time:         time,
			}
			responses = append(responses, res)
		}
	}
	return responses, nil
}

func (cu *commentUsecase) CreateComment(text string, uid, pid uint32) (commentResponse, error) {
	comment := &models.Comment{
		Text:   text,
		UserID: uid,
		PostID: pid,
	}

	err := validations.CommentValidate(comment)
	if err != nil {
		return commentResponse{}, err
	}

	err = cu.commentRepository.Save(comment)
	if err != nil {
		return commentResponse{}, err
	}

	commentedUser, imgURI, time, err := cu.GetUserData((*comment).UserID, (*comment).CreatedAt, (*comment).UpdatedAt)
	if err != nil {
		return commentResponse{}, err
	}
	res := commentResponse{
		ID:           (*comment).ID,
		Text:         (*comment).Text,
		UserID:       (*comment).UserID,
		UserNickname: commentedUser,
		UserImage:    imgURI,
		Time:         time,
	}

	return res, nil
}

func (cu *commentUsecase) UpdateComment(cid uint32, text string) (int64, error) {
	comment := &models.Comment{
		Text: text,
	}

	err := validations.CommentValidate(comment)
	if err != nil {
		return 0, err
	}

	rows, err := cu.commentRepository.Update(cid, comment)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (cu *commentUsecase) DeleteComment(cid uint32) error {
	err := cu.commentRepository.Delete(cid)
	if err != nil {
		return err
	}
	return nil
}

// ユーザー情報取得〜整形まで
func (cu *commentUsecase) GetUserData(uid uint32, createdAt, updatedAt time.Time) (string, string, string, error) {
	user, err := cu.userRepository.FindByID(uid)
	if err != nil {
		return "", "", "", err
	}

	imgURI, err := cu.GetUserImage(uid)
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
func (cu *commentUsecase) GetUserImage(uid uint32) (string, error) {
	img, err := cu.imageRepository.FindByUserID(uid)
	if err != nil {
		return "", err
	}

	err = cu.imageRepository.DownloadS3(img, "merpochi-users-image")
	if err != nil {
		return "", err
	}

	uri, err := security.Base64EncodeToString(img.Buf)
	if err != nil {
		return "", err
	}
	return uri, nil
}

type commentResponse struct {
	ID           uint32 `json:"id"`
	Text         string `json:"text"`
	UserID       uint32 `json:"user_id"`
	UserNickname string `json:"user_nickname"`
	UserImage    string `json:"user_image"`
	Time         string `json:"time"`
}
