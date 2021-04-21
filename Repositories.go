package main

import(
	"github.com/jinzhu/gorm"
)



type UserIRepository interface{
	Insert(user *User) (bool, error)
	GetAll() ([]User, error)	
}

type ProgressIRepository interface{
	Insert(progress *Progress) (bool, error)
	GetAll() ([]Progress, error)	
}

type TrainingIRepository interface{
	Insert(training *Training) (bool, error)
	GetAll() ([]Training, error)
	GetByGroupId(id int) ([]Training, error)
}

type TrainingGroupIRepository interface{
	Insert(trainingGroup *TrainingGroup) (bool, error)
	GetAll() ([]TrainingGroup, error)
}

type TrainRelatTrainGroupIRepository interface{
	Insert(trainRelatTG *TrainingRelationTrainingGroup) (bool, error)
	GetAll() ([]TrainingRelationTrainingGroup, error)
}


type UserRepository struct {
	db *gorm.DB
}

type ProgressRepository struct {
	db *gorm.DB
}

type TrainingRepository struct {
	db *gorm.DB
}

type TrainingGroupRepository struct {
	db *gorm.DB
}
type TrainRelatTrainGroupRepository struct {
	db *gorm.DB
}




func NewUserRepository() *UserRepository {
	return &UserRepository{GetDB()}
}


func NewProgressRepository() *ProgressRepository {
	return &ProgressRepository{GetDB()}
}

func NewTrainingRepository() *TrainingRepository {
	return &TrainingRepository{GetDB()}
}

func NewTrainingGroupRepository() *TrainingGroupRepository{
	return &TrainingGroupRepository{GetDB()}
}
func NewTrainRelatTrainGroupRepository() *TrainRelatTrainGroupRepository{
	return &TrainRelatTrainGroupRepository{GetDB()}
}



func (u UserRepository) Insert(user *User) (bool, error){
	res, err := user.Validate()
	if(!res){
		return false, err
	}
	err = u.db.Create(&user).Error
	if(err!=nil){
		return false, err
	}
	return true, nil
} 

func(u UserRepository) GetAll() ([]User, error){
	var users []User
	err := u.db.Find(&users).Error
	return users, err
}

func(p ProgressRepository) Insert(progress *Progress) (bool, error){
	res, err := progress.Validate()
	if(!res){
		return false, err
	}
	err = p.db.Create(&progress).Error
	if(err!=nil){
		return false, err
	}
	return true, nil
}

func(p ProgressRepository) GetAll() ([]Progress, error){
	var progresses []Progress
	err := p.db.Find(&progresses).Error
	return progresses, err
}


func(t TrainingRepository) Insert(training *Training) (bool, error){
	res, err :=training.Validate()
	if(!res){
		return false, err
	}
	err = t.db.Create(&training).Error
	if(err!=nil){
		return false, err
	}
	return true, nil
}

func(t TrainingRepository) GetAll() ([]Training, error){
	var trainings []Training
	err := t.db.Find(&trainings).Error
	return trainings, err
}
func (t TrainingRepository) GetByGroupId(id int) ([]Training, error){
	var res []Training
	 err := t.db.Joins("JOIN training_relation_training_groups ON training_relation_training_groups.training_id = trainings.id and training_relation_training_groups.training_group_id= ?", id).Find(&res).Error
	 if(err!=nil){
		 return res, err
	 }
	 return res, nil
}

func(t TrainingGroupRepository) Insert(trainingGroup *TrainingGroup) (bool, error){
	res, err :=trainingGroup.Validate()
	if(!res){
		return false, err
	}
	err = t.db.Create(&trainingGroup).Error
	if(err!=nil){
		return false, err
	}
	return true, nil
}

func(t TrainingGroupRepository) GetAll() ([]TrainingGroup, error){
	var trainingGroup []TrainingGroup
	err := t.db.Find(&trainingGroup).Error
	return trainingGroup, err
}

func(t TrainRelatTrainGroupRepository) Insert(trainingrelatTG *TrainingRelationTrainingGroup) (bool, error){
	res, err :=trainingrelatTG.Validate()
	if(!res){
		return false, err
	}
	err = t.db.Create(&trainingrelatTG).Error
	if(err!=nil){
		return false, err
	}
	return true, nil
}


func(t TrainRelatTrainGroupRepository) GetAll() ([]TrainingRelationTrainingGroup, error){
	var trainingrelatTG []TrainingRelationTrainingGroup
	err := t.db.Find(&trainingrelatTG).Error
	return trainingrelatTG, err
}
