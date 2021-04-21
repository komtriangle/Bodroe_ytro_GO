package main

import (
	"errors"
)


func(train *Training) Validate() (bool, error){
	if train.TrainingName == ""{
		return false, errors.New("Required TrainingName")
	}
	if train.Text == ""{
		return false, errors.New("Required Text")
	}
	if train.Photo == ""{
		return false, errors.New("Required Photo")
	}
	if train.Time == ""{
		return false, errors.New("Required Time")
	}
	return true, nil
}

func(trainingGroup *TrainingGroup) Validate() (bool, error){
	if trainingGroup.Name == ""{
		return false, errors.New("Required Name")
	}
	if trainingGroup.ShortDescription == ""{
		return false, errors.New("Required ShortDescription")
	}
	if trainingGroup.Description == ""{
		return false, errors.New("Required Description")
	}
	return true, nil
}

func(trainRelatTG *TrainingRelationTrainingGroup) Validate() (bool, error){
	if trainRelatTG.TrainingId <= 0 {
		return false, errors.New("Training id required must be greater than 0")
	}
	if trainRelatTG.TrainingGroupId <= 0{
		return false, errors.New("Training group id required and must be greater than 0")
	}
	return true, nil
}

func (user *User) Validate() (bool, error){
	if user.Name == "" {
		return false, errors.New("Name is required")
	}
	if user.Age <= 0{
		return false, errors.New("Age required and must be greater than 0")
	}
	return true, nil
}

func (progress *Progress) Validate() (bool, error){
	if progress.UserToken == "" {
		return false, errors.New("UserToken is required")
	}
	if progress.DateTime.IsZero(){
		return false, errors.New("DateTime is required")
	}
	return true, nil
}