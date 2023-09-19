package controller

import (
	controller "github.com/puranadhikari/go-boilerplate/internal/controller/impl"
	"github.com/puranadhikari/go-boilerplate/internal/service"
)

type Controllers struct {
	UserController IUserController
}

// WithControllers bootstrapping Controllers
func WithControllers(
	uc IUserController,
) *Controllers {
	return &Controllers{
		UserController: uc,
	}
}

// WithUserController bootstrapping user controller
func WithUserController(
	us service.IUserService,
) IUserController {
	return &controller.UserController{
		Us: us,
	}
}
