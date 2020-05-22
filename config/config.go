package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIPORT   = 0
	CLIENTURL = ""
)

// Load 環境変数の読み込み
func EnvLoad() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	APIPORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		APIPORT = 8000
	}

	CLIENTURL = os.Getenv("CLIENT_URL")
}
