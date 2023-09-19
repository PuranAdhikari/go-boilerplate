package repository

import (
	"github.com/puranadhikari/go-boilerplate/internal/model"
)

type IUserRepository interface {
	CreateUser(user *model.User) error
}
