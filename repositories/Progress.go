package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/komtriangle/Bodroe_ytro_GO/db"
)

type ProgressIRepository interface {
	Insert(progress *Progress) (bool, error)
	GetAll() ([]Progress, error)
}

type ProgressRepository struct {
	db *gorm.DB
}

func NewProgressRepository() *ProgressRepository {
	return &ProgressRepository{db.GetDB()}
}

func (p ProgressRepository) Insert(progress *Progress) (bool, error) {
	res, err := progress.Validate()
	if !res {
		return false, err
	}
	err = p.db.Create(&progress).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p ProgressRepository) GetAll() ([]Progress, error) {
	var progresses []Progress
	err := p.db.Find(&progresses).Error
	return progresses, err
}
