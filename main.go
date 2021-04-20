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


  router := mux.NewRouter()

  router.HandleFunc("/training", httpHandler.CreateTraining).Methods("POST")
  router.HandleFunc("/trainings", httpHandler.GetTrainings).Methods("GET")


  log.Fatal(http.ListenAndServe(":8080", router))
}
