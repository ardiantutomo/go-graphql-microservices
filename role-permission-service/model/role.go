package model

import "gorm.io/gorm"

// Role Database Model
type Role struct {
	gorm.Model
	RoleName  string `json:"role_name" gorm:"not null;index:role_name_unique,unique"`
	DeletedAt gorm.DeletedAt
}
