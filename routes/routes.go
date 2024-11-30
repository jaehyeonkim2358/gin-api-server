package routes

import (
	"example/gin-api-server/controllers/boards"
	"example/gin-api-server/database"
	httperrorhandler "example/gin-api-server/middleware/http_error_handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.Use(httperrorhandler.HandleError())
	api := router.Group("api")
	{
		client := database.GetClient()

		api.GET("/boards", func (ctx *gin.Context) {
			boards.Index(ctx, client)
		})
		api.GET("/boards/:id", func (ctx *gin.Context) {
			boards.Show(ctx, client)
		})
		api.POST("/boards", func (ctx *gin.Context) {
			boards.Create(ctx, client)
		})
		api.PATCH("/boards/:id", func (ctx *gin.Context) {
			boards.Update(ctx, client)
		})
	}
}
