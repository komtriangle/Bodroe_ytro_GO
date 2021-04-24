package repositories

import (
	"github.com/komtriangle/Bodroe_ytro_GO/db"
	"github.com/komtriangle/Bodroe_ytro_GO/models"
)

type TrainRelatTrainGroupIRepository interface {
	Insert(trainRelatTG *models.TrainingRelationTrainingGroup) (bool, error)
	GetAll() ([]models.TrainingRelationTrainingGroup, error)
}

type TrainRelatTrainGroupRepository struct {
	db db.Database
}

func NewTrainRelatTrainGroupRepository(database db.Database) *TrainRelatTrainGroupRepository {
	return &TrainRelatTrainGroupRepository{database}
}

func (t TrainRelatTrainGroupRepository) Insert(trainingrelatTG *models.TrainingRelationTrainingGroup) (bool, error) {
	res, err := trainingrelatTG.Validate()
	if !res {
		return false, err
	}
	res, err = t.db.CreateTrainRelatTG(trainingrelatTG)
	if !res {
		return false, err
	}
	return true, nil
}

func (t TrainRelatTrainGroupRepository) GetAll() ([]models.TrainingRelationTrainingGroup, error) {
	var trainingrelatTG []models.TrainingRelationTrainingGroup
	trainingrelatTG, err := t.db.GetAllTrainRelatTG()
	return trainingrelatTG, err
}
