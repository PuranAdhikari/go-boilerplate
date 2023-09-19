package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/puranadhikari/go-boilerplate/internal/controller"
)

func RegisterRoutes(r *gin.RouterGroup, ctl *controller.Controllers) {
	// Register auth routes
	AuthRoutes(r, ctl.UserController)
}
