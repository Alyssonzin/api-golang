package main

import (
	"api-golang/internal/container"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	pluggyController, userController := container.BuildContainer()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "HELLO WORLD",
		})
	})

	server.POST("/auth", pluggyController.CreatePluggyApikey)
	server.GET("/item/:item_id", pluggyController.GetItem)

	userGroup := server.Group("/user")
	userGroup.GET("/", userController.GetUsers)
	userGroup.POST("/", userController.CreateUser)
	userGroup.GET("/:id", userController.GetById)

	server.Run(":8080")
}
