package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/komtriangle/Bodroe_ytro_GO/db"
	"github.com/komtriangle/Bodroe_ytro_GO/repositories"
)

var err error

func main() {

	err = db.ConnectDB()
	if err != nil {
		panic(err)
	}
	MigrateDB()
	defer db.CloseDB()

	userRepo := repositories.NewUserRepository()
	progressRepo := repositories.NewProgressRepository()
	trainingRepo := repositories.NewTrainingRepository()
	trainingGroupRepo := repositories.NewTrainingGroupRepository()
	trainRelTG := repositories.NewTrainRelatTrainGroupRepository()
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

func MigrateDB() {
	var db *gorm.DB = db.GetDB()
	db.AutoMigrate(&repositories.Training{})
	db.AutoMigrate(&repositories.TrainingGroup{})
	db.AutoMigrate(&repositories.TrainingRelationTrainingGroup{})
	db.AutoMigrate(&repositories.User{})
	db.AutoMigrate(&repositories.Progress{})

	db.Model(&repositories.TrainingRelationTrainingGroup{}).AddForeignKey("Training_Id", "Trainings(Id)", "RESTRICT", "RESTRICT")
	db.Model(&repositories.TrainingRelationTrainingGroup{}).AddForeignKey("Training_Group_Id", "training_groups(Id)", "RESTRICT", "RESTRICT")
	db.Model(&repositories.Progress{}).AddForeignKey("User_Token", "users(Id)", "RESTRICT", "RESTRICT")

}
