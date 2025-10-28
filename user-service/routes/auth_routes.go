package routes

import (
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	r := router.Group("/auth")
	{
		r.POST("/register", controllers.Register)
		r.POST("/login", controllers.Login)
	}
}
