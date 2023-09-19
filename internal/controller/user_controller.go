package controller

import (
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Register(ctx *gin.Context)
}
