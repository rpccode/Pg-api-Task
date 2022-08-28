package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null;unique_index" `
	Description string `json:"description"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      uint   `json:"user_id"`
}
