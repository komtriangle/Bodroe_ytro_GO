package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/komtriangle/Bodroe_ytro_GO/db"
	"github.com/komtriangle/Bodroe_ytro_GO/models"
	"github.com/komtriangle/Bodroe_ytro_GO/repositories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	database_url := os.Getenv("DATABASE_URL")
	dialect := os.Getenv("DIALECT")
	PostgresDB, err := gorm.Open(dialect, database_url)
	if err != nil {
		panic(err)
	}
	database := db.DatabaseGorm{DB: PostgresDB}
	if err != nil {
		panic(err)
	}
	fmt.Println(database.GetDB())
	MigrateDB(database)
	defer database.CloseDB()

	userRepo := repositories.NewUserRepository(database)
	progressRepo := repositories.NewProgressRepository(database)
	trainingRepo := repositories.NewTrainingRepository(database)
	trainingGroupRepo := repositories.NewTrainingGroupRepository(database)
	trainRelTG := repositories.NewTrainRelatTrainGroupRepository(database)
	httpHandler := HttpHandler{userRepo, progressRepo, trainingRepo, trainingGroupRepo, trainRelTG}

	router := mux.NewRouter()

	router.HandleFunc("/Training", httpHandler.CreateTraining).Methods("POST")
	router.HandleFunc("/Trainings", httpHandler.GetTrainings).Methods("GET")
	router.HandleFunc("/TrainingGroup", httpHandler.CreateTrainingGroup).Methods("POST")
	router.HandleFunc("/TrainingGroups", httpHandler.GetAllTrainingGroups).Methods("GET")
	router.HandleFunc("/TrainingRelationTrainingGroup", httpHandler.CreateTrainRelatTrainGroup).Methods("POST")
	router.HandleFunc("/TrainingsFromGroup/{id:[0-9]+}", httpHandler.GetTrainingsFromGroup).Methods("GET")
	router.HandleFunc("/User", httpHandler.CreateUser).Methods("POST")
	router.HandleFunc("/Users", httpHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/Progress", httpHandler.CreateProgress).Methods("POST")
	router.HandleFunc("/Progresses", httpHandler.GetAllProgresses).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func MigrateDB(database db.Database) {
	var db *gorm.DB = database.GetDB()
	db.AutoMigrate(&models.Training{})
	db.AutoMigrate(&models.TrainingGroup{})
	db.AutoMigrate(&models.TrainingRelationTrainingGroup{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Progress{})

	db.Model(&models.TrainingRelationTrainingGroup{}).AddForeignKey("Training_Id", "Trainings(Id)", "RESTRICT", "RESTRICT")
	db.Model(&models.TrainingRelationTrainingGroup{}).AddForeignKey("Training_Group_Id", "training_groups(Id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Progress{}).AddForeignKey("User_Token", "users(Id)", "RESTRICT", "RESTRICT")

}
