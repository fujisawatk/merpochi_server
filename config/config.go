package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// グローバル変数
var (
	ENV       = ""
	APIPORT   = 0
	CLIENTURL = ""
	DBDRIVER  = ""
	DBURL     = ""
	SECRETKEY []byte
	AWSID     = ""
	AWSKEY    = ""
)

// EnvLoad 環境変数の読み込み
func EnvLoad() {
	var err error

	ENV = os.Getenv("ENV")
	// 本番環境は環境変数管理をSSMで行うため、.env読み込み不可。
	if "development" == ENV {
		err = godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	APIPORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	CLIENTURL = os.Getenv("CLIENT_URL")

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	SECRETKEY = []byte(os.Getenv("API_SECRET"))

	AWSID = os.Getenv("AWS_ACCESS_KEY_ID")
	AWSKEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
}
