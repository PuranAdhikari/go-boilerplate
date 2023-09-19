package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/puranadhikari/go-boilerplate/internal/dto"
	"github.com/puranadhikari/go-boilerplate/internal/service"
	"github.com/puranadhikari/go-boilerplate/internal/util"
)

type UserController struct {
	Us service.IUserService
}

func (c *UserController) Signup(ctx *gin.Context) {
	var input dto.SignupRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.RespondWithError(ctx, 422, err)
		return
	}

	response, err := c.Us.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(200, response)
}
