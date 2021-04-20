package main

import (
	"fmt"
	"log"
  "os"
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
var handler Handler


func main(){
  
  database := os.Getenv("DATABASE_URL")
  dialect := os.Getenv("DIALECT")
  db, err = gorm.Open(dialect, database)

  
  repo := Handler{db}
  httpHandler := HttpHandler{repo}
  
  if(err!=nil){
    log.Fatal(err)
  }else{
    fmt.Println("Success")
  }

  defer db.Close()

  db.Exec("PRAGMA foreign_keys = ON")
  db.AutoMigrate(&Training{})
  db.AutoMigrate(&TrainingGroup{})
  db.AutoMigrate(&TrainingRelationTrainingGroup{})

  db.Model(&TrainingRelationTrainingGroup{}).AddForeignKey("Training_Id", "Trainings(Id)", "RESTRICT", "RESTRICT")
  db.Model(&TrainingRelationTrainingGroup{}).AddForeignKey("Training_Group_Id", "training_groups(Id)", "RESTRICT", "RESTRICT")

  
  router := mux.NewRouter()

  router.HandleFunc("/Training", httpHandler.CreateTraining).Methods("POST")
  router.HandleFunc("/Trainings", httpHandler.GetTrainings).Methods("GET")
  router.HandleFunc("/TrainingGroup", httpHandler.CreateTrainingGroup).Methods("POST")
  router.HandleFunc("/TrainingGroups", httpHandler.GetAllTrainingGroups).Methods("GET")
  router.HandleFunc("/TrainingRelationTrainingGroup", httpHandler.CreateTrainRelatTrainGroup).Methods("POST")


  log.Fatal(http.ListenAndServe(":8080", router))
}
