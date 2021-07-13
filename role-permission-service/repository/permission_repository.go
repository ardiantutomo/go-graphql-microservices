package repository

import (
	"errors"
	"log"
	"role-permission-service/model"

	"gorm.io/gorm"
)

type permissionRepository struct {
	DB *gorm.DB
}

// PermissionRepository : represent the permission's repository contract
type PermissionRepository interface {
	Save(model.Permission) (model.Permission, error)
	GetAll() ([]model.Permission, error)
	WithTrx(*gorm.DB) permissionRepository
	Migrate() error
}

// NewPermissionRepository -> returns new permission repository
func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return permissionRepository{
		DB: db,
	}
}

func (u permissionRepository) Migrate() error {
	log.Print("[PermissionRepository]...Migrate")
	return u.DB.AutoMigrate(&model.Permission{})
}

func (u permissionRepository) Save(permission model.Permission) (model.Permission, error) {
	log.Print("[PermissionRepository]...Save")
	err := u.DB.Create(&permission).Error
	if err != nil {
		err = errors.New("Permission creation failed")
	}
	return permission, err

}

func (u permissionRepository) GetAll() (permissions []model.Permission, err error) {
	log.Print("[PermissionRepository]...Get All")
	err = u.DB.Find(&permissions).Error
	return permissions, err
}

func (u permissionRepository) WithTrx(trxHandle *gorm.DB) permissionRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}
