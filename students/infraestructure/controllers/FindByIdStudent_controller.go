package controllers

import (
	"api-hexagonal/students/application"
	"api-hexagonal/students/infraestructure/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FindByIdController struct {
	useCase *application.FindById
}

func NewFindByIdController(useCase *application.FindById) *FindByIdController {
	return &FindByIdController{useCase: useCase}
}

func (fbi_c *FindByIdController) Run(ctx *gin.Context) {
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

	student, err := fbi_c.useCase.Run(idInt)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	response := responses.NewResponseStudentFound(student)

	ctx.JSON(http.StatusOK, response)

}