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
	HandleImageCreate(w http.ResponseWriter, r *http.Request)
	HandleImageGet(w http.ResponseWriter, r *http.Request)
	HandleImageUpdate(w http.ResponseWriter, r *http.Request)
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

// HandleUserImageCreate ユーザー情報を登録
func (ih imageHandler) HandleImageCreate(w http.ResponseWriter, r *http.Request) {
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

	img, err := ih.imageUsecase.CreateImage(uint32(uid), file)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, img)
}

// HandleUserImageUploadAndCreate ユーザー情報を登録
func (ih imageHandler) HandleImageGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	imgURI, err := ih.imageUsecase.GetImage(uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, imgURI)
}

// HandleUserImageUpdate ユーザー画像を更新
func (ih imageHandler) HandleImageUpdate(w http.ResponseWriter, r *http.Request) {
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

	rows, err := ih.imageUsecase.UpdateImage(uint32(uid), file)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, rows)
}
