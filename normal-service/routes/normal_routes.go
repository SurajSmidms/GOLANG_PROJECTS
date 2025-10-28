package routes

import (
	"net/http"
	"normal-service/middleware"

	"github.com/gin-gonic/gin"
)

func NormalRoutes(router *gin.Engine) {
	r := router.Group("/normal")
	r.Use(middleware.AuthMiddleware())
	{
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello from Normal Service!"})
		})
	}
}
