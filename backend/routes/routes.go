package routes

import (
	"Platform/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api") // Prefix all routes with "/api"
	{
		api.POST("/register/user", controllers.RegisterUser)
	}
}
