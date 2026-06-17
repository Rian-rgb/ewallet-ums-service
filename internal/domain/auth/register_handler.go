package auth

import "github.com/gin-gonic/gin"

type IRegisterHandler interface {
	Register(ctx *gin.Context)
}
