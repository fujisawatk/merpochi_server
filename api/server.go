package api

import (
	"fmt"
	"log"
	"merpochi_server/api/router"
	"merpochi_server/config"
	"net/http"
)

// Run サーバ起動
func Run() {
	config.EnvLoad()
	fmt.Printf("\n\tListening [::]:%d\n", config.APIPORT)
	listen(config.APIPORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), (r)))
}
