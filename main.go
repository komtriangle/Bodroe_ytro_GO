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

  db.AutoMigrate(&Training{})
  db.AutoMigrate(&TrainingGroup{})


  router := mux.NewRouter()

  router.HandleFunc("/Training", httpHandler.CreateTraining).Methods("POST")
  router.HandleFunc("/Trainings", httpHandler.GetTrainings).Methods("GET")
  router.HandleFunc("/TrainingGroup", httpHandler.CreateTrainingGroup).Methods("POST")
  router.HandleFunc("/TrainingGroups", httpHandler.GetAllTrainingGroups).Methods("GET")


  log.Fatal(http.ListenAndServe(":8080", router))
}
