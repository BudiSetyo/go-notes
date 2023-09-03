package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UserID      uint   `json:"user_id"`
	CategoryID  uint   `json:"category_id"`
}
