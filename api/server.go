package api

import (
	"fmt"
	"log"
	"merpochi_server/api/auto"
	"merpochi_server/config"
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/router"
	"merpochi_server/interfaces/healthcheck"
	"net/http"
)

// Run サーバ起動
func Run() {
	config.EnvLoad()
	// 開発環境では起動毎にDBテーブル、テストデータリセット
	if config.ENV == "development" {
		auto.TestLoad()
	}
	fmt.Printf("\n\tListening [::]:%d\n", config.APIPORT)
	listen(config.APIPORT)
}

func listen(port int) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	r := router.New()
	r.HandleFunc("/health", healthcheck.HandleHealthCheck)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), (r)))
}
