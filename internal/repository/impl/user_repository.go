package repository

import (
	"github.com/puranadhikari/go-boilerplate/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}
