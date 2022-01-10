package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateRouter(router *gin.Engine) *gin.Engine {
	routes := router.Group("/api/v1/")
	{
		routes.GET("/user/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
