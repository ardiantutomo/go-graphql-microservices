package repository

import (
	"errors"
	"log"
	"role-permission-service/model"

	"gorm.io/gorm"
)

type roleRepository struct {
	DB *gorm.DB
}

// RoleRepository : represent the role's repository contract
type RoleRepository interface {
	Save(model.Role) (model.Role, error)
	GetAll() ([]model.Role, error)
	WithTrx(*gorm.DB) roleRepository
	Migrate() error
}

// NewRoleRepository -> returns new role repository
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepository{
		DB: db,
	}
}

func (u roleRepository) Migrate() error {
	log.Print("[RoleRepository]...Migrate")
	return u.DB.AutoMigrate(&model.Role{})
}

func (u roleRepository) Save(role model.Role) (model.Role, error) {
	log.Print("[RoleRepository]...Save")
	err := u.DB.Create(&role).Error
	if err != nil {
		err = errors.New("Role creation failed")
	}
	return role, err

}

func (u roleRepository) GetAll() (roles []model.Role, err error) {
	log.Print("[RoleRepository]...Get All")
	err = u.DB.Find(&roles).Error
	return roles, err
}

func (u roleRepository) WithTrx(trxHandle *gorm.DB) roleRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}
