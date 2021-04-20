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