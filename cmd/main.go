package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/puranadhikari/go-boilerplate/config"
	"github.com/puranadhikari/go-boilerplate/database"
	"github.com/puranadhikari/go-boilerplate/internal/controller"
	"github.com/puranadhikari/go-boilerplate/internal/repository"
	"github.com/puranadhikari/go-boilerplate/internal/routes"
	"github.com/puranadhikari/go-boilerplate/internal/service"
	"github.com/puranadhikari/go-boilerplate/internal/util"
)

func main() {
	cfg := config.InitializeConfig()
	db := database.Connect(cfg.DB)
	database.AutoMigrate(db)

	// Compile time dependency injection
	ur := repository.WithUserRepository(db)
	us := service.WithUserService(ur)
	uc := controller.WithUserController(us)

	ctl := controller.WithControllers(uc)
	r := gin.Default()

	// Define the global route group
	routes.RegisterRoutes(util.RoutePrefix, r, ctl)

	// Start the Gin server.
	r.Run(fmt.Sprintf(":%d", cfg.App.Port))
}
