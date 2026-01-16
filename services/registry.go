package services

import (
	"user-service/repositories"
	services "user-service/services/user"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
}

// GetUser implements IServiceRegistry.
func (r *Registry) GetUser() services.IUserService {
	return services.NewUserService(r.repository)
}

type IServiceRegistry interface {
	GetUser() services.IUserService
}

func NewServiceRegistry(repo repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{
		repository: repo,
	}
}
