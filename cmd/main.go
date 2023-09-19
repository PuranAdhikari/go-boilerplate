package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/puranadhikari/go-boilerplate/config"
	"github.com/puranadhikari/go-boilerplate/internal/controller"
	"github.com/puranadhikari/go-boilerplate/internal/model"
	"github.com/puranadhikari/go-boilerplate/internal/repository"
	"github.com/puranadhikari/go-boilerplate/internal/routes"
	"github.com/puranadhikari/go-boilerplate/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "dev", "Environment (dev|prod|stg)")
	flag.Parse()
	// Load the configuration.
	cfg, err := config.LoadConfig("./config", env)
	if err != nil {
		panic(fmt.Sprintf("Failed to load %s configuration: %v", env, err))
	}

	// Use cfg for accessing configurations.
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=disable port=%d",
		cfg.DB.Host, cfg.DB.User, cfg.DB.DBName, cfg.DB.Password, cfg.DB.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema.
	db.AutoMigrate(&model.User{})

	// Compile time dependency injection
	ur := repository.WithUserRepository(db)
	us := service.WithUserService(ur)
	uc := controller.WithUserController(us)

	ctl := controller.WithControllers(uc)
	r := gin.Default()

	// Define the global route group
	v1 := r.Group("/api/v1")
	// Set up routes.
	routes.RegisterRoutes(v1, ctl)

	// Start the Gin server.
	r.Run(fmt.Sprintf(":%d", cfg.App.Port))
}
