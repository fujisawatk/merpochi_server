package database

import (
	"merpochi_server/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB    *gorm.DB
	DbErr error
)

// Connect DB接続
func Connect() (*gorm.DB, error) {
	DB, DbErr = gorm.Open(config.DBDRIVER, config.DBURL)
	if DbErr != nil {
		return nil, DbErr
	}
	return DB, nil
}
