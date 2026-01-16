package controllers

import (
	controllers "user-service/controllers/user"
	"user-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IUserControllerRegistry interface {
	GetUserController() controllers.IUserController
}

func NewControllerRegistry(service services.IServiceRegistry) IUserControllerRegistry {
	return &Registry{
		service: service,
	}
}

// GetUserController implements IUserControllerRegistry.
func (r *Registry) GetUserController() controllers.IUserController {
	return controllers.NewUserController(r.service)
}