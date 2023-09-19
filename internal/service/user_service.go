package service

import (
	"github.com/puranadhikari/go-boilerplate/internal/dto"
)

type IUserService interface {
	Signup(email, password string) (*dto.SignUpResponse, error)
}
