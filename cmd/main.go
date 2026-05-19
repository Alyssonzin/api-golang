package main

import (
	"api-golang/controller"
	"api-golang/db"
	"api-golang/internal/integrations/pluggy"
	"api-golang/repository"
	"api-golang/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	pluggyClient := pluggy.NewPluggyClient("http://localhost:8000")

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

	server.GET("/test", func(ctx *gin.Context) {

		data, err := pluggyClient.GetItem(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Data(http.StatusOK, "application/json", data)
	})

	UserGroup := server.Group("/user")
	UserGroup.GET("/", UserController.GetUsers)
	UserGroup.POST("/", UserController.CreateUser)
	UserGroup.GET("/:id", UserController.GetById)

	server.Run(":8080")
}
