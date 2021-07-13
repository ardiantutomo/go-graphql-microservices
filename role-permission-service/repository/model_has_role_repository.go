package repository

import (
	"errors"
	"log"
	"role-permission-service/model"

	"gorm.io/gorm"
)

type modelHasRoleRepository struct {
	DB *gorm.DB
}

// ModelHasRoleRepository : represent the modelHasRole's repository contract
type ModelHasRoleRepository interface {
	Save(model.ModelHasRole) (model.ModelHasRole, error)
	GetAll() ([]model.ModelHasRole, error)
	WithTrx(*gorm.DB) modelHasRoleRepository
	Migrate() error
}

// NewModelHasRoleRepository -> returns new modelHasRole repository
func NewModelHasRoleRepository(db *gorm.DB) ModelHasRoleRepository {
	return modelHasRoleRepository{
		DB: db,
	}
}

func (u modelHasRoleRepository) Migrate() error {
	log.Print("[ModelHasRoleRepository]...Migrate")
	return u.DB.AutoMigrate(&model.ModelHasRole{})
}

func (u modelHasRoleRepository) Save(modelHasRole model.ModelHasRole) (model.ModelHasRole, error) {
	log.Print("[ModelHasRoleRepository]...Save")
	err := u.DB.Create(&modelHasRole).Error
	if err != nil {
		err = errors.New("modelHasRole creation failed")
	}
	return modelHasRole, err

}

func (u modelHasRoleRepository) GetAll() (modelHasRoles []model.ModelHasRole, err error) {
	log.Print("[ModelHasRoleRepository]...Get All")
	err = u.DB.Find(&modelHasRoles).Error
	return modelHasRoles, err
}

func (u modelHasRoleRepository) WithTrx(trxHandle *gorm.DB) modelHasRoleRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}
