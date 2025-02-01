package controllers

import (
	"api-hexagonal/career/application"
	"api-hexagonal/career/domain/entities"
	"api-hexagonal/career/infraestructure/responses"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCareerController struct {
	useCase *application.CreateCareer
}

func NewCreateCareerController(useCase *application.CreateCareer) *CreateCareerController {
	return &CreateCareerController{useCase: useCase}
}

func (cc_c *CreateCareerController) Run(ctx *gin.Context) {
	var career entities.Career

	if err := ctx.ShouldBindJSON(&career); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if career.Name == "" && career.Duration <= 0 && career.Type == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Los campos están vacíos o son inválidos"})
		return
	}

	id, err := cc_c.useCase.Run(&career)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// Definimos los encabezados
	ctx.Writer.Header().Set("Content-Type", "application/vnd.api+json")

	ctx.Writer.Header().Set("Location", fmt.Sprintf("http://localhost:8080/careers/%d", id))

	// Reasignamos el id de la carrera
	career.Id = id

	response := responses.NewResponseCareerCreated(&career)

	ctx.JSON(http.StatusCreated, response)

}