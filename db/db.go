package db

import (
	"code.aliyun.com/microstack/rpc-user/db/models"
	myGorm "code.aliyun.com/microstack/utils/db/gorm"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Migrate() error {
	if DB == nil {
		return errors.New("migrate: db is not connected")
	}

	if err := DB.AutoMigrate(models.All...).Error; err != nil {
		return err
	}
	return nil
}

func open() {
	db, err := myGorm.Open(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"))
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	DB = db

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(1000)
	if os.Getenv("RUN_MODE") == "dev" {
		DB.LogMode(true)
	}
}

func init() {
	open()
}
