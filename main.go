package main

import (
	"fmt"
	"log"
  "os"
	"net/http"
  "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
  "./Repositories"
)


func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello")
  }


var db *gorm.DB
var err error
var handler Handler

func GetDB() *gorm.DB{
  return db
}


func main(){
  
  database := os.Getenv("DATABASE_URL")
  dialect := os.Getenv("DIALECT")
  db, err = gorm.Open(dialect, database)

  
  repo := Handler{db}
  //userRepo := NewUserRepository()
  httpHandler := HttpHandler{repo}
  
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
  //router.HandleFunc("/User", httpHandler.CreateUser).Methods("POST")
  


  log.Fatal(http.ListenAndServe(":8080", router))
}


