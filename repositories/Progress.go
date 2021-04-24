package repositories

import (
	"github.com/komtriangle/Bodroe_ytro_GO/db"
	"github.com/komtriangle/Bodroe_ytro_GO/models"
)

type ProgressIRepository interface {
	Insert(progress *models.Progress) (bool, error)
	GetAll() ([]models.Progress, error)
}

type ProgressRepository struct {
	db db.Database
}

func NewProgressRepository(database db.Database) *ProgressRepository {
	return &ProgressRepository{database}
}

func (p ProgressRepository) Insert(progress *models.Progress) (bool, error) {
	res, err := progress.Validate()
	if !res {
		return false, err
	}
	res, err = p.db.CreateProgress(progress)
	if !res {
		return false, err
	}
	return true, nil
}

func (p ProgressRepository) GetAll() ([]models.Progress, error) {
	var progresses []models.Progress
	progresses, err := p.db.GetAllProgresses()
	return progresses, err
}
