package main

import (
	"github.com/jinzhu/gorm"
)

type Repository interface{
	InsertTranining(training *Training) 
	GetAllTrainings() ([]Training, error)
	InsertTrainingGroup(trainingGroup *TrainingGroup) (bool, error)
	GetAllTrainingGroups() ([]TrainingGroup, error)
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

func (h Handler) GetAllTrainings() ([]Training, error){
	var training []Training
	err := h.db.Find(&training).Error
	return training, err
}

func(h Handler) InsertTrainingGroup(trainingGroup *TrainingGroup) (bool, error){
	res, err :=trainingGroup.Validate()
	if(!res){
		return res, err
	}
	results := h.db.Create(&trainingGroup)
	if(results.Error != nil){
		return res, results.Error
	}
	return true, nil
}

func(h Handler) GetAllTrainingGroups() ([]TrainingGroup, error){
	var trainingGroups []TrainingGroup
	err = h.db.Find(&trainingGroups).Error

	return trainingGroups, err
}

