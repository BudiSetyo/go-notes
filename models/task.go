package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
	UserID      uint   `json:"user_id" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}
