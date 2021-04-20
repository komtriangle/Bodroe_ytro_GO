package main

import (
	"github.com/jinzhu/gorm"
)

type Handler struct {
	db *gorm.DB
}


func CreateHandler(db_ *gorm.DB) *Handler{
	return &Handler{db: db_}
}