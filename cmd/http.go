package cmd

import (
	_ "ewallet-ums/docs"
	"ewallet-ums/infra"
	"github.com/Rian-rgb/ewallet-common-lib/config"
	"github.com/Rian-rgb/ewallet-common-lib/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

// @title           E Wallet API (User Management Service)
// @version         1.0
// @description     Dokumentasi API Brilian.
// @host            localhost:8080
// @BasePath        /api/v1
func ServeHTTP(dependency *infra.Dependency, appDeps *infra.AppDependencies) {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/user")
		{
			users.POST("/register", dependency.RegisterAPI.Register)
			users.POST("/login", dependency.LoginAPI.Login)
			users.DELETE("/logout", middleware.AuthMiddleware(appDeps.JWTManager.ValidateToken, *appDeps.RedisRepo), dependency.LogoutAPI.Logout)
			users.PUT("/refresh-token", middleware.RefreshTokenMiddleware(appDeps.JWTManager.ValidateToken, *appDeps.RedisRepo), dependency.RefreshTokenAPI.RefreshToken)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := config.GetEnv("PORT", "8080")
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
