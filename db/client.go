package db

import (
	"github.com/jinzhu/gorm"
	g "github.com/micro-stacks/utils/db/gorm"
	"log"
	"os"
)

var Client *gorm.DB

func init() {
	db, err := g.Open(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"))
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	Client = db

	Client.DB().SetMaxIdleConns(10)
	Client.DB().SetMaxOpenConns(1000)
	if os.Getenv("RUN_MODE") == "dev" {
		Client.LogMode(true)
	}
}
