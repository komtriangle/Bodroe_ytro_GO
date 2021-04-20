package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	db *gorm.DB
}


func CreateHandler(db_ *gorm.DB) *Handler{
	return &Handler{db: db_}
}

func (h *Handler) CreateTraining(w http.ResponseWriter, r *http.Request){
	var training Training
	json.NewDecoder(r.Body).Decode(&training)
	fmt.Println(training)
	h.db.Create(&training)
}

func (h *Handler) GetTrainings(w http.ResponseWriter, r *http.Request){
	var training []Training
	h.db.Find(&training)
	json.NewEncoder(w).Encode(&training)
}