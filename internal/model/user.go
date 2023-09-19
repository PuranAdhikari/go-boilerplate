package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // This provides fields like ID, CreatedAt, UpdatedAt, etc.
	Email      string `gorm:"type:varchar(100);uniqueIndex"`
	Password   string `gorm:"size:255"` // store hashed password
}
