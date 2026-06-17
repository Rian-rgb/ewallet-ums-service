package http

import (
	"ewallet-ums/infra"

	"github.com/Rian-rgb/ewallet-common-lib/middleware"
	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(
	api *gin.RouterGroup,
	dependency *infra.Dependency,
	appDeps *infra.AppDependencies,
) {
	auth := api.Group("/auth")

	auth.POST(
		"/login",
		dependency.LoginAPI.Login,
	)

	auth.PUT(
		"/refresh-token",
		middleware.RefreshTokenMiddleware(
			appDeps.JWTManager.ValidateToken,
			*appDeps.RedisRepo,
		),
		dependency.RefreshTokenAPI.RefreshToken,
	)

	auth.DELETE(
		"/logout",
		middleware.AuthMiddleware(
			appDeps.JWTManager.ValidateToken,
			*appDeps.RedisRepo,
		),
		dependency.LogoutAPI.Logout,
	)

	auth.POST(
		"/register",
		dependency.RegisterAPI.Register,
	)
}
