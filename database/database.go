package database

import (
	"fmt"
	"github.com/puranadhikari/go-boilerplate/config"
	"github.com/puranadhikari/go-boilerplate/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dc config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=disable port=%d",
		dc.Host, dc.User, dc.DBName, dc.Password, dc.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	// Migrate the schema.
	db.AutoMigrate(&model.User{})
}
