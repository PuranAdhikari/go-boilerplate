package repository

import (
	repository "github.com/puranadhikari/go-boilerplate/internal/repository/impl"
	"gorm.io/gorm"
)

// WithUserRepository bootstrapping user repository
func WithUserRepository(
	db *gorm.DB,
) IUserRepository {
	return &repository.UserRepository{
		DB: db,
	}
}
