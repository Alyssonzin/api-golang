package controller

import (
	"api-golang/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(uc usecase.UserUseCase) userController {
	return userController{
		userUseCase: uc,
	}
}

func (u *userController) GetUsers(ctx *gin.Context) {
	users, err := u.userUseCase.GetUsers()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)
}
