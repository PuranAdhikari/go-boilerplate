package service

import (
	"github.com/puranadhikari/go-boilerplate/internal/repository"
	service "github.com/puranadhikari/go-boilerplate/internal/service/impl"
)

// WithUserService bootstrapping user service
func WithUserService(
	ur repository.IUserRepository,
) IUserService {
	return &service.UserService{
		Ur: ur,
	}
}
