package model

import "gorm.io/gorm"

// Role Database Model
type RoleHasPermission struct {
	gorm.Model
	RoleId       int
	PermissionId int
	Permission   Permission
	Role         Role
	DeletedAt    gorm.DeletedAt
}
