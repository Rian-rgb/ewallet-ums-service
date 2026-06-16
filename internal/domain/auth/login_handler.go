package auth

import "github.com/gin-gonic/gin"

type ILoginHandler interface {
	Login(c *gin.Context)
}
