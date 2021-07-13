package repository

import (
	"errors"
	"log"
	"role-permission-service/model"

	"gorm.io/gorm"
)

type modelHasPermissionRepository struct {
	DB *gorm.DB
}

// ModelHasPermissionRepository : represent the permission's repository contract
type ModelHasPermissionRepository interface {
	Save(model.ModelHasPermission) (model.ModelHasPermission, error)
	GetAll() ([]model.ModelHasPermission, error)
	WithTrx(*gorm.DB) modelHasPermissionRepository
	Migrate() error
}

// NewModelHasPermissionRepository -> returns new permission repository
func NewModelHasPermissionRepository(db *gorm.DB) ModelHasPermissionRepository {
	return modelHasPermissionRepository{
		DB: db,
	}
}

func (u modelHasPermissionRepository) Migrate() error {
	log.Print("[ModelHasPermissionRepository]...Migrate")
	return u.DB.AutoMigrate(&model.ModelHasPermission{})
}

func (u modelHasPermissionRepository) Save(permission model.ModelHasPermission) (model.ModelHasPermission, error) {
	log.Print("[ModelHasPermissionRepository]...Save")
	err := u.DB.Create(&permission).Error
	if err != nil {
		err = errors.New("ModelHasPermission creation failed")
	}
	return permission, err

}

func (u modelHasPermissionRepository) GetAll() (modelHasPermissions []model.ModelHasPermission, err error) {
	log.Print("[ModelHasPermissionRepository]...Get All")
	err = u.DB.Find(&modelHasPermissions).Error
	return modelHasPermissions, err
}

func (u modelHasPermissionRepository) WithTrx(trxHandle *gorm.DB) modelHasPermissionRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}
