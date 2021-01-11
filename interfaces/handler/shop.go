package handler

import (
	"encoding/json"
	"io/ioutil"
	"merpochi_server/domain/models"
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
)

// ShopHandler Shopに対するHandlerのインターフェイス
type ShopHandler interface {
	HandleShopsSearch(w http.ResponseWriter, r *http.Request)
	HandleShopCreate(w http.ResponseWriter, r *http.Request)
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

// HandleShopsGet 外部APIで取得した店舗が登録されているか検索
func (sh *shopHandler) HandleShopsSearch(w http.ResponseWriter, r *http.Request) {
	// body形式 → ["XX0000", ...]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody shopSearchRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	counts, err := sh.shopUsecase.SearchShops(requestBody.ShopCodes, requestBody.UserID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, counts)
}

// HandleShopCreate 店舗情報を登録
func (sh shopHandler) HandleShopCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	var requestBody models.Shop
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	shop, err := sh.shopUsecase.CreateShop(requestBody)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, shop)
}

type shopSearchRequest struct {
	ShopCodes []string `json:"shop_codes"`
	UserID    uint32   `json:"user_id"`
}
