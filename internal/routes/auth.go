package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/puranadhikari/go-boilerplate/internal/controller"
)

func AuthRoutes(r *gin.RouterGroup, userController controller.IUserController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		// Add logout route, refresh token routes, etc. here as needed.
	}
}
