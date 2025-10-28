package main

import (
	"normal-service/config"
	"normal-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	routes.NormalRoutes(r)

	r.Run(":8083")
}
