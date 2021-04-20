package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
  "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Training struct{
    Id int  `gorm:"primaryKey"`
    TrainingName string
    Text string
    Photo string
    Time string
}


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


  handler := CreateHandler(db)
  
  if(err!=nil){
    log.Fatal(err)
  }else{
    fmt.Println("Success")
  }

  defer db.Close()

  db.AutoMigrate(&Training{})


  router := mux.NewRouter()

  router.HandleFunc("/training", handler.CreateTraining).Methods("POST")
  router.HandleFunc("/trainings", handler.GetTrainings).Methods("GET")


  log.Fatal(http.ListenAndServe(":8080", router))
}
