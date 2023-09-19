package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/puranadhikari/go-boilerplate/internal/dto"
	"github.com/puranadhikari/go-boilerplate/internal/service"
	"github.com/puranadhikari/go-boilerplate/internal/util"
	"github.com/puranadhikari/go-boilerplate/pkg/logger"
	"net/http"
)

type UserController struct {
	Us service.IUserService
}

func (c *UserController) Register(ctx *gin.Context) {
	var input dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Log.With(err).Errorf("error while validating user register request")
		ctx.JSON(http.StatusUnprocessableEntity, util.HandleValidationError(err))
		return
	}

	response, err := c.Us.Register(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, response)
}
