package model

import "gorm.io/gorm"

// Role Database Model
type ModelHasPermission struct {
	gorm.Model
	PermissionId int
	ModelId      int
	Permission   Permission
	DeletedAt    gorm.DeletedAt
}
