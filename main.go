package main

import (
	"fmt"
	"log"
	"net/http"
  "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello")
  }


var db *gorm.DB
var err error

func GetDB() *gorm.DB{
  return db
}


func main(){
  
  //database := os.Getenv("DATABASE_URL")
  //dialect := os.Getenv("DIALECT")
  db, err = gorm.Open("postgres", "postgresql://postgres:1234@localhost/bodroe_ytro_GO?sslmode=disable")

  
  userRepo := NewUserRepository()
  progressRepo := NewProgressRepository()
  trainingRepo := NewTrainingRepository()
  trainingGroupRepo := NewTrainingGroupRepository()
  trainRelTG :=NewTrainRelatTrainGroupRepository()
  httpHandler := HttpHandler{userRepo, progressRepo, trainingRepo, trainingGroupRepo,trainRelTG}

  
  if(err!=nil){
    log.Fatal(err)
  }else{
    fmt.Println("Success")
  }

  defer db.Close()
  db.AutoMigrate(&Training{})
  db.AutoMigrate(&TrainingGroup{})
  db.AutoMigrate(&TrainingRelationTrainingGroup{})
  db.AutoMigrate(&User{})
  db.AutoMigrate(&Progress{})

  db.Model(&TrainingRelationTrainingGroup{}).AddForeignKey("Training_Id", "Trainings(Id)", "RESTRICT", "RESTRICT")
  db.Model(&TrainingRelationTrainingGroup{}).AddForeignKey("Training_Group_Id", "training_groups(Id)", "RESTRICT", "RESTRICT")
  db.Model(&Progress{}).AddForeignKey("User_Token", "users(Id)", "RESTRICT", "RESTRICT")

  
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


