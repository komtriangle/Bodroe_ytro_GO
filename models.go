package main


type Training struct{
    Id int  `gorm:"primaryKey"`
    TrainingName string
    Text string
    Photo string
    Time string
}

type TrainingGroup struct{
	Id int  `gorm:"primaryKey"`
	Name string
	ShortDescription string
	Description string
}
