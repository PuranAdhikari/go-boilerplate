package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/puranadhikari/go-boilerplate/internal/controller"
)

func RegisterRoutes(prefix string, r *gin.Engine, ctl *controller.Controllers) {
	grp := r.Group(prefix)
	AuthRoutes(grp, ctl.UserController)
}
