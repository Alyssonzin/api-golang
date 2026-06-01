package container

import (
	"api-golang/config"
	"api-golang/controller"
	"api-golang/db"
	"api-golang/internal/integrations/pluggy"
	"api-golang/repository"
	"api-golang/usecase"
)

func BuildContainer() (*controller.PluggyController, *controller.UserController) {
	envs := config.LoadEnvs(".env")

	pluggyClient := pluggy.NewPluggyClient(envs.PluggyClientID, envs.PluggyClientSecret)
	pluggyUseCase := usecase.NewPluggyUseCase(pluggyClient)
	pluggyController := controller.NewPluggyController(pluggyUseCase)

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(dbConnection)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)

	return pluggyController, userController
}
