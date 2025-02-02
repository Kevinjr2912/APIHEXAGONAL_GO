package controllers

import (
	"api-hexagonal/career/application"
	"api-hexagonal/career/domain/entities"
	"api-hexagonal/career/infraestructure/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateCareerController struct {
	useCase *application.UpdateCareer
}

func NewUpdateCareerController(useCase *application.UpdateCareer) *UpdateCareerController {
	return &UpdateCareerController{useCase: useCase}
}

func (uc_c *UpdateCareerController) Run(ctx *gin.Context) {
	idStr, exists := ctx.Params.Get("id")
	var career entities.Career

	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id no proporcionado"})
		return
	}

	idInt, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id no válido"})
		return
	}

	if err := ctx.ShouldBindJSON(&career); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if career.Name == "" && career.Duration <= 0 && career.Type == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Los campos están vacíos o son inválidos"})
		return
	}

	err = uc_c.useCase.Run(idInt, &career)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	response := responses.NewResponseCareerUpdated(idInt, &career)

	ctx.JSON(http.StatusOK, response)

}