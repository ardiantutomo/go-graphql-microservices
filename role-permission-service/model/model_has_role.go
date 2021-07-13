package model

import "gorm.io/gorm"

// Role Database Model
type ModelHasRole struct {
	gorm.Model
	RoleId    int
	ModelId   int
	Role      Role
	DeletedAt gorm.DeletedAt
}
