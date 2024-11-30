package main

import (
	"github.com/jaehyeonkim2358/gin-api-server/database"
	"github.com/jaehyeonkim2358/gin-api-server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	defer database.GetClient().Close()

	router := gin.Default()
	routes.InitRoutes(router)
	router.Run("localhost:8888")
}
