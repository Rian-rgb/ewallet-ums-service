package auth

import "github.com/gin-gonic/gin"

type ILogoutHandler interface {
	Logout(ctx *gin.Context)
}
