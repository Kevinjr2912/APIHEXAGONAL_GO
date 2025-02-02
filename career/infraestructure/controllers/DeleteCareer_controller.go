package controllers

import (
	"api-hexagonal/career/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCareerController struct {
	useCase *application.DeleteCareer
}

func NewDeleteCareerController(useCase *application.DeleteCareer) *DeleteCareerController {
	return &DeleteCareerController{useCase: useCase}
}

func (dc_c *DeleteCareerController) Run(ctx *gin.Context) {
	idStr, exists := ctx.Params.Get("id")
	
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id no proporcionado"})
		return
	}

	idInt, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id inválido"})
		return
	}

	err = dc_c.useCase.Run(idInt)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}