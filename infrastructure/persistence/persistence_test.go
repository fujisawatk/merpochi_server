package persistence

import (
	"fmt"
	"log"
	"merpochi_server/domain/models"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	// プール作成
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// パラメーター設定
	opts := dockertest.RunOptions{
		Repository:   "mysql",
		Tag:          "5.7",
		Env:          []string{"MYSQL_ROOT_PASSWORD=password"},
		ExposedPorts: []string{"3306"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306": {
				{HostIP: "0.0.0.0", HostPort: "3307"},
			},
		},
	}

	// コンテナ実行
	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// DB接続
	if err := pool.Retry(func() error {
		databaseConnStr := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", "root",
			"password", "tcp(0.0.0.0:3307)", "mysql")
		db, err = gorm.Open("mysql", databaseConnStr)
		if err != nil {
			log.Println("Database not ready yet (it is booting up, wait for a few tries)...")
			return err
		}

		// 疎通確認
		return db.DB().Ping()
	}); err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// 初期設定
	log.Println("Initialize test database...")
	initTestDatabase()

	log.Println("Run the actual test cases...")
	code := m.Run()

	// コンテナ削除
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func initTestDatabase() {
	// マイグレーション
	ccommandAndArgs := []string{
		"-path", "../../db", "-env", "test", "up",
	}
	out, err := exec.Command("goose", ccommandAndArgs...).Output()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(out))

	// 初期データ挿入
	for _, user := range users {
		err = db.Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, shop := range shops {
		err = db.Model(&models.Shop{}).Create(&shop).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, favorite := range favorites {
		err = db.Model(&models.Favorite{}).Create(&favorite).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, comment := range comments {
		err = db.Model(&models.Comment{}).Create(&comment).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, station := range stations {
		err = db.Model(&models.Station{}).Create(&station).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}
