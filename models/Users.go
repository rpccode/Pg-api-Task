package models
import ("gorm.io/gorm")
type User struct {
	gorm.Model

	firstName string
	lastName string
	email string
	Tasks []Task
}