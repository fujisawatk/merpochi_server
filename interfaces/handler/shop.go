package handler

import (
	"encoding/json"
	"io/ioutil"
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ShopHandler Shopに対するHandlerのインターフェイス
type ShopHandler interface {
	HandleShopsGet(w http.ResponseWriter, r *http.Request)
	HandleShopCreate(w http.ResponseWriter, r *http.Request)
	HandleShopGet(w http.ResponseWriter, r *http.Request)
}

type shopHandler struct {
	shopUsecase usecase.ShopUsecase
}

// NewShopHandler Shopデータに関するHandlerを生成
func NewShopHandler(us usecase.ShopUsecase) ShopHandler {
	return &shopHandler{
		shopUsecase: us,
	}
}

// HandleShopsGet 店舗情報を全件取得
func (sh shopHandler) HandleShopsGet(w http.ResponseWriter, r *http.Request) {
	shops, err := sh.shopUsecase.GetShops()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, shops)
}

// HandleShopCreate 店舗情報を登録
func (sh shopHandler) HandleShopCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody shopCreateRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	shop, err := sh.shopUsecase.CreateShop(requestBody.Code)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, shop)
}

// HandleShopGet 店舗情報を1件取得
func (sh shopHandler) HandleShopGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	shop, err := sh.shopUsecase.GetShop(uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, shop)
}

type shopCreateRequest struct {
	Code string `json:"code"`
}
