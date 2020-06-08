package handler

import (
	"merpochi_server/interfaces/responses"
	"merpochi_server/usecase"
	"net/http"
)

// ShopHandler Shopに対するHandlerのインターフェイス
type ShopHandler interface {
	HandleShopsGet(w http.ResponseWriter, r *http.Request)
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

// HandleShopsGet ユーザー情報を全件取得
func (sh shopHandler) HandleShopsGet(w http.ResponseWriter, r *http.Request) {
	shops, err := sh.shopUsecase.GetShops()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, shops)
}
