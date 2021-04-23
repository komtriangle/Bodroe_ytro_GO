package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/komtriangle/Bodroe_ytro_GO/db"
)

type TrainingGroupIRepository interface {
	Insert(trainingGroup *TrainingGroup) (bool, error)
	GetAll() ([]TrainingGroup, error)
}

type TrainingGroupRepository struct {
	db *gorm.DB
}

func NewTrainingGroupRepository() *TrainingGroupRepository {
	return &TrainingGroupRepository{db.GetDB()}
}

func (t TrainingGroupRepository) Insert(trainingGroup *TrainingGroup) (bool, error) {
	res, err := trainingGroup.Validate()
	if !res {
		return false, err
	}
	err = t.db.Create(&trainingGroup).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t TrainingGroupRepository) GetAll() ([]TrainingGroup, error) {
	var trainingGroup []TrainingGroup
	err := t.db.Find(&trainingGroup).Error
	return trainingGroup, err
}
