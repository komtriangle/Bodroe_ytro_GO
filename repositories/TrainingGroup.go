package repositories

import (
	"github.com/komtriangle/Bodroe_ytro_GO/db"
	"github.com/komtriangle/Bodroe_ytro_GO/models"
)

type TrainingGroupIRepository interface {
	Insert(trainingGroup *models.TrainingGroup) (bool, error)
	GetAll() ([]models.TrainingGroup, error)
}

type TrainingGroupRepository struct {
	db db.Database
}

func NewTrainingGroupRepository(database db.Database) *TrainingGroupRepository {
	return &TrainingGroupRepository{database}
}

func (t TrainingGroupRepository) Insert(trainingGroup *models.TrainingGroup) (bool, error) {
	res, err := trainingGroup.Validate()
	if !res {
		return false, err
	}
	res, err = t.db.CreateTrainingGroup(trainingGroup)
	if !res {
		return false, err
	}
	return true, nil
}

func (t TrainingGroupRepository) GetAll() ([]models.TrainingGroup, error) {
	var trainingGroup []models.TrainingGroup
	trainingGroup, err := t.db.GetAllTrainingGroups()
	return trainingGroup, err
}
