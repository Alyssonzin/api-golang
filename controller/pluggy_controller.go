package controller

import (
	"api-golang/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PluggyController struct {
	pluggyApikeyUseCase usecase.PluggyUseCase
}

func NewPluggyController(pluggyApikeyUseCase usecase.PluggyUseCase) *PluggyController {
	return &PluggyController{pluggyApikeyUseCase: pluggyApikeyUseCase}
}

func (c *PluggyController) CreatePluggyApikey(ctx *gin.Context) {
	apiKey, err := c.pluggyApikeyUseCase.CreatePluggyApikey(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"api_key": apiKey})
}

func (c *PluggyController) GetItem(ctx *gin.Context) {
	itemID := ctx.Param("item_id")
	item, err := c.pluggyApikeyUseCase.GetItem(ctx, itemID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, item)
}
