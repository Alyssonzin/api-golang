package main

import (
	"api-golang/controller"
	"api-golang/db"
	"api-golang/repository"
	"api-golang/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, error := db.ConnectDB()

	if error != nil {
		panic(error)
	}

	UserRepository := repository.NewUserRepository(dbConnection)

	UserUseCase := usecase.NewUserUseCase(UserRepository)

	UserController := controller.NewUserController(UserUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "HELLO WORLD",
		})
	})

	UserGroup := server.Group("/user")
	UserGroup.GET("/", UserController.GetUsers)
	UserGroup.POST("/", UserController.CreateUser)

	server.Run(":8080")
}
