package handler

import (
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ImageHandler Userに対するHandlerのインターフェイス
type ImageHandler interface {
	HandleImageUpload(w http.ResponseWriter, r *http.Request)
}

type imageHandler struct {
	imageUsecase usecase.ImageUsecase
}

// NewImageHandler Imageデータに関するHandlerを生成
func NewImageHandler(iu usecase.ImageUsecase) ImageHandler {
	return &imageHandler{
		imageUsecase: iu,
	}
}

// HandleUserImageUploadAndCreate ユーザー情報を登録
func (ih imageHandler) HandleImageUpload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	img, err := ih.imageUsecase.UploadImage(uint32(uid), file)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, img)
}
