package controller

import (
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Signup(ctx *gin.Context)
}
