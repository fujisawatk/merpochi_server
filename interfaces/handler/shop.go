package handler

import (
	"encoding/json"
	"io/ioutil"
	"merpochi_server/domain/models"
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
	HandleShopFavoritesGet(w http.ResponseWriter, r *http.Request)
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

// HandleShopsGet 外部APIで取得した各店舗に紐付く情報（コメント数）を取得
func (sh shopHandler) HandleShopsGet(w http.ResponseWriter, r *http.Request) {
	// body形式 → ["XX0000", ...]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var requestBody []string
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	counts, err := sh.shopUsecase.GetShops(requestBody)
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

// HandleShopFavoriteGet 店舗情報を1件取得
func (sh shopHandler) HandleShopFavoritesGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	favorites, err := sh.shopUsecase.GetShopFavorites(uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, favorites)
}
