package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func GetDB() *gorm.DB {
	return db
}

func ConnectDB() {
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}
	database := os.Getenv("DATABASE_URL")
	dialect := os.Getenv("DIALECT")
	db, err = gorm.Open(dialect, database)

	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	db.Close()
}
