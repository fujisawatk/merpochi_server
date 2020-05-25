package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIPORT   = 0
	CLIENTURL = ""
	DBDRIVER  = ""
	DBURL     = ""
)

// EnvLoad 環境変数の読み込み
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

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_PROT"), os.Getenv("DB_NAME"))
}
