package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/komtriangle/Bodroe_ytro_GO/db"
)

type TrainRelatTrainGroupIRepository interface {
	Insert(trainRelatTG *TrainingRelationTrainingGroup) (bool, error)
	GetAll() ([]TrainingRelationTrainingGroup, error)
}

type TrainRelatTrainGroupRepository struct {
	db *gorm.DB
}

func NewTrainRelatTrainGroupRepository() *TrainRelatTrainGroupRepository {
	return &TrainRelatTrainGroupRepository{db.GetDB()}
}

func (t TrainRelatTrainGroupRepository) Insert(trainingrelatTG *TrainingRelationTrainingGroup) (bool, error) {
	res, err := trainingrelatTG.Validate()
	if !res {
		return false, err
	}
	err = t.db.Create(&trainingrelatTG).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t TrainRelatTrainGroupRepository) GetAll() ([]TrainingRelationTrainingGroup, error) {
	var trainingrelatTG []TrainingRelationTrainingGroup
	err := t.db.Find(&trainingrelatTG).Error
	return trainingrelatTG, err
}
