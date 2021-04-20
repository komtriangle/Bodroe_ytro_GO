package main

import (
	"github.com/jinzhu/gorm"
)

type Repository interface{
	InsertTranining(training *Training) 
	GetAllTrainings() []Training
}

type Handler struct {
	db *gorm.DB
}

func CreateHandler(db_ *gorm.DB) *Handler{
	return &Handler{db: db_}
}

func (h Handler) InsertTranining(training *Training){
	res, err := training.Validate()
	if(!res){
		panic(err)
	}
    result := h.db.Create(&training)
	
	if(result.Error!=nil){
		panic(result.Error)
	}
}

func (h Handler) GetAllTrainings() []Training{
	var training []Training
	h.db.Find(&training)
	return training
}

