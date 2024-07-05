package models

import "gorm.io/gorm"

type TodoUnit struct {
	gorm.Model
	Name        string `json:"name"`
	IsCompleted bool   `gorm:"default:false"`
}
