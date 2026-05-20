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

	pluggyClient := pluggy.NewPluggyClient()

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

	server.POST("/auth", func(ctx *gin.Context) {

		data, err := pluggyClient.CreateApiKey(ctx.Request.Context(), "client_id", "client_secret")
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		resp := struct {
			ApiKey string `json:"apiKey"`
		}{
			ApiKey: data.ApiKey,
		}

		ctx.JSON(http.StatusOK, resp)
	})

	UserGroup := server.Group("/user")
	UserGroup.GET("/", UserController.GetUsers)
	UserGroup.POST("/", UserController.CreateUser)
	UserGroup.GET("/:id", UserController.GetById)

	server.Run(":8080")
}
