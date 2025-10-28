package routes

import (
	"book-service/controllers"
	"book-service/middleware"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	r := router.Group("/books")
	r.Use(middleware.AuthMiddleware())
	{
		r.POST("/", controllers.CreateBook)
		r.GET("/", controllers.GetBooks)
	}
}
