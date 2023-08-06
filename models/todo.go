package models

import "github.com/jinzhu/gorm"

type ToDo struct {
	gorm.Model
	Task string `json:"task" grom:"type:text not null"`
	Done bool   `json:"done" gorm:"not null"`
}
