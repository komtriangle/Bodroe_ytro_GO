package repositories

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

func InitDB() {
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
	db.AutoMigrate(&Training{})
	db.AutoMigrate(&TrainingGroup{})
	db.AutoMigrate(&TrainingRelationTrainingGroup{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Progress{})

	db.Model(&TrainingRelationTrainingGroup{}).AddForeignKey("Training_Id", "Trainings(Id)", "RESTRICT", "RESTRICT")
	db.Model(&TrainingRelationTrainingGroup{}).AddForeignKey("Training_Group_Id", "training_groups(Id)", "RESTRICT", "RESTRICT")
	db.Model(&Progress{}).AddForeignKey("User_Token", "users(Id)", "RESTRICT", "RESTRICT")
}

func CloseDB() {
	db.Close()
}
