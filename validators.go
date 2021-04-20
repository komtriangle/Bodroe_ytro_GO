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