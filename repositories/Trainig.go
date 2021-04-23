package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/komtriangle/Bodroe_ytro_GO/db"
)

type TrainingIRepository interface {
	Insert(training *Training) (bool, error)
	GetAll() ([]Training, error)
	GetByGroupId(id int) ([]Training, error)
}

type TrainingRepository struct {
	db *gorm.DB
}

func NewTrainingRepository() *TrainingRepository {
	return &TrainingRepository{db.GetDB()}
}

func (t TrainingRepository) Insert(training *Training) (bool, error) {
	res, err := training.Validate()
	if !res {
		return false, err
	}
	err = t.db.Create(&training).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t TrainingRepository) GetAll() ([]Training, error) {
	var trainings []Training
	err := t.db.Find(&trainings).Error
	return trainings, err
}
func (t TrainingRepository) GetByGroupId(id int) ([]Training, error) {
	var res []Training
	err := t.db.Joins("JOIN training_relation_training_groups ON training_relation_training_groups.training_id = trainings.id and training_relation_training_groups.training_group_id= ?", id).Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
