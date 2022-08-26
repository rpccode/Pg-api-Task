package models

import (
	"gorm.io/gorm"
)
type User struct {
	gorm.Model

	firstName string `gorm: "not null"`
	lastName  string `gorm: "not null"`
	email     string `gorm: "not null; unique_index"`
	Tasks     []Task
}