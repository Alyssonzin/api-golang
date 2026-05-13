package controller

import (
	"api-golang/dto"
	"api-golang/model"
	"api-golang/usecase"
	"fmt"
	"net/http"
	"strconv"

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

func (u *userController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := dto.Response{
			Message: "ID do usuário é obrigatório",
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		res := dto.Response{
			Message: "ID do usuário precisa ser número",
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	user, err := u.userUseCase.GetById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		res := dto.Response{
			Message: "Usuário não encontrado",
		}
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *userController) CreateUser(ctx *gin.Context) {
	var user model.User

	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	newUser, err := u.userUseCase.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newUser)
}
