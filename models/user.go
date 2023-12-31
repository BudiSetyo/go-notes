package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;uniqueIndex;not null" json:"username" validate:"required"`
	Email    string `gorm:"unique;uniqueIndex;not null" json:"email" validate:"required"`
	Password string `gorm:"not null" json:"password" validate:"required"`
}
