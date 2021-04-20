package main

import (
  "encoding/json"
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

func create_training(w http.ResponseWriter, r *http.Request){
	var training Training
	json.NewDecoder(r.Body).Decode(&training)
  
  
	fmt.Println(training)
	db.Create(&training)
  }
  
  func get_trainings(w http.ResponseWriter, r *http.Request){
	  var training []Training
	
	  db.Find(&training)
	
	  json.NewEncoder(w).Encode(&training)
  }
  

var db *gorm.DB
var err error


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

  router.HandleFunc("/training", create_training).Methods("POST")
  router.HandleFunc("/trainings", get_trainings).Methods("GET")


  log.Fatal(http.ListenAndServe(":8080", router))
}
