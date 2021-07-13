package repository

import (
	"errors"
	"log"
	"role-permission-service/model"

	"gorm.io/gorm"
)

type roleHasPermissionRepository struct {
	DB *gorm.DB
}

// RoleHasPermissionRepository : represent the roleHasPermission's repository contract
type RoleHasPermissionRepository interface {
	Save(model.RoleHasPermission) (model.RoleHasPermission, error)
	GetAll() ([]model.RoleHasPermission, error)
	WithTrx(*gorm.DB) roleHasPermissionRepository
	Migrate() error
}

// NewRoleHasPermissionRepository -> returns new roleHasPermission repository
func NewRoleHasPermissionRepository(db *gorm.DB) RoleHasPermissionRepository {
	return roleHasPermissionRepository{
		DB: db,
	}
}

func (u roleHasPermissionRepository) Migrate() error {
	log.Print("[RoleHasPermissionRepository]...Migrate")
	return u.DB.AutoMigrate(&model.RoleHasPermission{})
}

func (u roleHasPermissionRepository) Save(roleHasPermission model.RoleHasPermission) (model.RoleHasPermission, error) {
	log.Print("[RoleHasPermissionRepository]...Save")
	err := u.DB.Create(&roleHasPermission).Error
	if err != nil {
		err = errors.New("RoleHasPermission creation failed")
	}
	return roleHasPermission, err

}

func (u roleHasPermissionRepository) GetAll() (roleHasPermissions []model.RoleHasPermission, err error) {
	log.Print("[RoleHasPermissionRepository]...Get All")
	err = u.DB.Find(&roleHasPermissions).Error
	return roleHasPermissions, err
}

func (u roleHasPermissionRepository) WithTrx(trxHandle *gorm.DB) roleHasPermissionRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}
