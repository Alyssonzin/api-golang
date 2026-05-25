package controller

import (
	"api-golang/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pluggyController struct {
	pluggyApikeyUseCase usecase.PluggyApikeyUseCase
}

func NewPluggyController(pluggyApikeyUseCase usecase.PluggyApikeyUseCase) pluggyController {
	return pluggyController{pluggyApikeyUseCase: pluggyApikeyUseCase}
}

func (c *pluggyController) CreatePluggyApikey(ctx *gin.Context) {
	apiKey, err := c.pluggyApikeyUseCase.CreatePluggyApikey(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"api_key": apiKey})
}
