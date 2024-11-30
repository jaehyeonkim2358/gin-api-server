package main

import (
	"example/gin-api-server/database"
	"example/gin-api-server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	defer database.GetClient().Close()

	router := gin.Default()
	routes.InitRoutes(router)
	router.Run("localhost:8888")
}
