package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/komtriangle/Bodroe_ytro_GO/models"
)

type Database interface {
	CreateTraining(training *models.Training) (bool, error)
	GetAllTrainings() ([]models.Training, error)
	GetTrainingByGroupId(id int) ([]models.Training, error)
	CreateTrainingGroup(trainingGroup *models.TrainingGroup) (bool, error)
	GetAllTrainingGroups() ([]models.TrainingGroup, error)
	CreateTrainRelatTG(trainRelatTG *models.TrainingRelationTrainingGroup) (bool, error)
	GetAllTrainRelatTG() ([]models.TrainingRelationTrainingGroup, error)
	CreateUser(user *models.User) (bool, error)
	GetAllUsers() ([]models.User, error)
	CreateProgress(progress *models.Progress) (bool, error)
	GetAllProgresses() ([]models.Progress, error)
	CloseDB()
	GetDB() *gorm.DB
}

type DatabaseGorm struct {
	DB *gorm.DB
}

func (d DatabaseGorm) GetDB() *gorm.DB {
	return d.DB
}

func (d DatabaseGorm) CloseDB() {
	d.DB.Close()
}

func (d DatabaseGorm) CreateTraining(training *models.Training) (bool, error) {
	err := d.DB.Create(&training).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d DatabaseGorm) GetAllTrainings() ([]models.Training, error) {
	var trainings []models.Training
	err := d.DB.Find(&trainings).Error
	if err != nil {
		return nil, err
	}
	return trainings, err
}

func (d DatabaseGorm) GetTrainingByGroupId(id int) ([]models.Training, error) {
	var trainings []models.Training
	err := d.DB.Joins("JOIN training_relation_training_groups ON training_relation_training_groups.training_id = trainings.id and training_relation_training_groups.training_group_id= ?", id).Find(&trainings).Error
	if err != nil {
		return nil, err
	}
	return trainings, err
}

func (d DatabaseGorm) CreateTrainingGroup(trainingGroup *models.TrainingGroup) (bool, error) {
	err := d.DB.Create(&trainingGroup).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d DatabaseGorm) GetAllTrainingGroups() ([]models.TrainingGroup, error) {
	var trainingGroup []models.TrainingGroup
	err := d.DB.Find(&trainingGroup).Error
	return trainingGroup, err
}

func (d DatabaseGorm) CreateTrainRelatTG(trainingrelatTG *models.TrainingRelationTrainingGroup) (bool, error) {
	err := d.DB.Create(&trainingrelatTG).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func (d DatabaseGorm) GetAllTrainRelatTG() ([]models.TrainingRelationTrainingGroup, error) {
	var trainingrelatTG []models.TrainingRelationTrainingGroup
	err := d.DB.Find(&trainingrelatTG).Error
	return trainingrelatTG, err
}

func (d DatabaseGorm) CreateUser(user *models.User) (bool, error) {
	err := d.DB.Create(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func (d DatabaseGorm) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := d.DB.Find(&users).Error
	return users, err
}

func (d DatabaseGorm) CreateProgress(progress *models.Progress) (bool, error) {
	err := d.DB.Create(&progress).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func (d DatabaseGorm) GetAllProgresses() ([]models.Progress, error) {
	var progresses []models.Progress
	err := d.DB.Find(&progresses).Error
	return progresses, err
}
