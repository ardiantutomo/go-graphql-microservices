package repository

import (
	"auth-service/model"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// UserRepository : represent the user's repository contract
type UserRepository interface {
	Save(model.User) (model.User, error)
	GetAll() ([]model.User, error)
	WithTrx(*gorm.DB) userRepository
	Migrate() error
}

// NewUserRepository -> returns new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&model.User{})
}

func (u userRepository) Save(user model.User) (model.User, error) {
	log.Print("[UserRepository]...Save")
	err := u.DB.Create(&user).Error
	if err != nil {
		err = errors.New("Email is already registered")
	}
	return user, err

}

func (u userRepository) GetAll() (users []model.User, err error) {
	log.Print("[UserRepository]...Get All")
	err = u.DB.Find(&users).Error
	return users, err

}

func (u userRepository) WithTrx(trxHandle *gorm.DB) userRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}
