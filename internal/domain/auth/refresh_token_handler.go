package auth

import "github.com/gin-gonic/gin"

type IRefreshTokenHandler interface {
	RefreshToken(c *gin.Context)
}
