package main

import (
	"api-golang/controller"
	"api-golang/db"
	"api-golang/internal/integrations/pluggy"
	"api-golang/repository"
	"api-golang/usecase"
	"net/http"

	"api-golang/config"

	"github.com/gin-gonic/gin"
)

func main() {
	envs := config.LoadEnvs(".env")
	server := gin.Default()

	pluggyClient := pluggy.NewPluggyClient(envs.PluggyClientID, envs.PluggyClientSecret)
	pluggyApikeyUseCase := usecase.NewPluggyApikeyUseCase(pluggyClient)
	PluggyController := controller.NewPluggyController(pluggyApikeyUseCase)

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

	server.POST("/auth", PluggyController.CreatePluggyApikey)

	UserGroup := server.Group("/user")
	UserGroup.GET("/", UserController.GetUsers)
	UserGroup.POST("/", UserController.CreateUser)
	UserGroup.GET("/:id", UserController.GetById)

	server.Run(":8080")
}
