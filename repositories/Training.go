package repositories

import (
	"github.com/komtriangle/Bodroe_ytro_GO/db"
	"github.com/komtriangle/Bodroe_ytro_GO/models"
)

type TrainingIRepository interface {
	Insert(training *models.Training) (bool, error)
	GetAll() ([]models.Training, error)
	GetByGroupId(id int) ([]models.Training, error)
}

type TrainingRepository struct {
	db db.Database
}

func NewTrainingRepository(database db.Database) *TrainingRepository {
	return &TrainingRepository{database}
}

func (t TrainingRepository) Insert(training *models.Training) (bool, error) {
	res, err := training.Validate()
	if !res {
		return false, err
	}
	res, err = t.db.CreateTraining(training)
	if !res {
		return false, err
	}
	return true, nil
}

func (t TrainingRepository) GetAll() ([]models.Training, error) {
	var trainings []models.Training
	trainings, err := t.db.GetAllTrainings()
	return trainings, err
}
func (t TrainingRepository) GetByGroupId(id int) ([]models.Training, error) {
	var trainings []models.Training
	trainings, err := t.db.GetTrainingByGroupId(id)
	if err != nil {
		return trainings, err
	}
	return trainings, nil
}
