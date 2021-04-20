package main


type Training struct{
    Id int  `gorm:"primaryKey"`
    TrainingName string
    Text string
    Photo string
    Time string
	TrainRelatTG []TrainingRelationTrainingGroup `gorm:"FOREIGNKEY:Training_Id;ASSOCIATION_FOREIGNKEY:id"`
}

type TrainingGroup struct{
	Id int  `gorm:"primaryKey"`
	Name string
	ShortDescription string
	Description string
	TrainRelatTG []TrainingRelationTrainingGroup `gorm:"FOREIGNKEY:Training_Group_Id;ASSOCIATION_FOREIGNKEY:id"`
}

type TrainingRelationTrainingGroup struct {
	Id int  `gorm:"primaryKey"`
	TrainingId int
	TrainingGroupId int 
}
