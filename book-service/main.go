package main

import (
	"book-service/config"
	"book-service/models"
	"book-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Book{})

	r := gin.Default()
	routes.BookRoutes(r)

	r.Run(":8082")
}
