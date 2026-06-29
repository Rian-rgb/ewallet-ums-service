package http

import (
	"ewallet-ums/infra"
	"github.com/Rian-rgb/ewallet-common-lib/config"

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
		dependency.LoginHdl.Login,
	)

	auth.PUT(
		"/refresh-token",
		middleware.RefreshTokenMiddleware(
			appDeps.JWTManager.ValidateToken,
			*appDeps.RedisRepo,
			config.GetEnv("SECRET_KEY_ENCRYPT", ""),
		),
		dependency.RefreshTokenHdl.RefreshToken,
	)

	auth.DELETE(
		"/logout",
		middleware.AuthMiddleware(
			appDeps.JWTManager.ValidateToken,
			*appDeps.RedisRepo,
		),
		dependency.LogoutHdl.Logout,
	)

	auth.POST(
		"/register",
		dependency.RegisterHdl.Register,
	)
}
