package model

import "gorm.io/gorm"

// User Database Model
type User struct {
	gorm.Model
	Fullname  string `json:"fullname" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	DeletedAt gorm.DeletedAt
}
