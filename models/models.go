package models

import (
	"time"
)

type Training struct {
	Id           int `gorm:"primaryKey autoIncrement"`
	TrainingName string
	Text         string
	Photo        string
	Time         string
	TrainRelatTG []TrainingRelationTrainingGroup `gorm:"FOREIGNKEY:Training_Id;ASSOCIATION_FOREIGNKEY:id"`
}

type TrainingGroup struct {
	Id               int `gorm:"primaryKey autoIncrement"`
	Name             string
	ShortDescription string
	Description      string
	TrainRelatTG     []TrainingRelationTrainingGroup `gorm:"FOREIGNKEY:Training_Group_Id;ASSOCIATION_FOREIGNKEY:id"`
}

type TrainingRelationTrainingGroup struct {
	Id              int `gorm:"primaryKey autoIncrement"`
	TrainingId      int
	TrainingGroupId int
}

type User struct {
	Id       string `gorm:"primaryKey autoIncrement"`
	Name     string
	Age      int
	Progress []Progress `gorm:"FOREIGNKEY:User_Token;ASSOCIATION_FOREIGNKEY:id"`
}

type Progress struct {
	Id        int `gorm:"primaryKey autoIncrement"`
	UserToken string
	DateTime  time.Time
}
