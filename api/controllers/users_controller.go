package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"merpochi_server/api/database"
	"merpochi_server/api/models"
	"net/http"
)

// CreateUser ユーザー新規登録
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// リクエスト情報を読み込む
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	// モデルの初期化
	user := models.User{}
	// モデルに保管
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
	}
	// データベースに接続（Gorm使用）
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// データベースに保存
	err = db.Debug().Model(&models.User{}).Create(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	// 保存後の処理
	fmt.Println("成功")
}
