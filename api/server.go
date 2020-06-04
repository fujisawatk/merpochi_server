package api

import (
	"fmt"
	"log"
	"merpochi_server/api/auto"
	"merpochi_server/config"
	"merpochi_server/infrastructure/database"
	"merpochi_server/infrastructure/router"
	"net/http"
)

// Run サーバ起動
func Run() {
	config.EnvLoad()
	auto.TestLoad()
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
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), (r)))
}
