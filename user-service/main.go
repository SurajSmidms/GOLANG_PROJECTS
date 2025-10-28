package main

import (
	"user-service/config"
	"user-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	routes.AuthRoutes(r)
	r.Run(":8081")
}
