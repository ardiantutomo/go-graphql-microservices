package service

import (
	"role-permission-service/model"
	"role-permission-service/repository"
)

// PermissionService : represent the user's service contract
type PermissionService interface {
	Save(model.Permission) (model.Permission, error)
	GetAll() ([]model.Permission, error)
}

type permissionService struct {
	permissionRepository repository.PermissionRepository
}

// NewPermissionService -> returns new user service
func NewPermissionService(r repository.PermissionRepository) PermissionService {
	return permissionService{
		permissionRepository: r,
	}
}

// fungsi untuk ke repositori
func (u permissionService) Save(user model.Permission) (model.Permission, error) {
	return u.permissionRepository.Save(user)
}

func (u permissionService) GetAll() ([]model.Permission, error) {

	return u.permissionRepository.GetAll()
}
