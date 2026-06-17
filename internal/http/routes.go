package http

import (
	"ewallet-ums/infra"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	dependency *infra.Dependency,
	appDeps *infra.AppDependencies,
) {
	api := router.Group("/api/v1")

	registerAuthRoutes(
		api,
		dependency,
		appDeps,
	)

	registerSwaggerRoutes(router)
}
