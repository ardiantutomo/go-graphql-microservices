package model

import "gorm.io/gorm"

// Permission Database Model
type Permission struct {
	gorm.Model
	PermissionName string `json:"permission_name" gorm:"not null;index:permission_name_unique,unique"`
	DeletedAt      gorm.DeletedAt
}
