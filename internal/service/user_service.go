package service

import (
	"github.com/puranadhikari/go-boilerplate/internal/dto"
)

type IUserService interface {
	Register(email, password string) (*dto.RegisterResponse, error)
}
