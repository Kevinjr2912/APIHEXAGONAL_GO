package controllers

import (
	"api-hexagonal/students/application"
	"api-hexagonal/students/domain/entities"
	"api-hexagonal/students/infraestructure/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateStudentController struct {
	useCase *application.UpdateStudent
}

func NewUpdateStudentController(useCase *application.UpdateStudent) *UpdateStudentController {
	return &UpdateStudentController{useCase: useCase}
}

func (us_c *UpdateStudentController) Run(ctx *gin.Context) {
	idStr, exists := ctx.Params.Get("id")

	var student entities.Student

	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id no proporcionado"})
		return
	}

	idInt, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id inválido"})
		return
	}

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if student.Name == "" && student.Age <= 0 && student.PhoneNumber <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Los campos están vacíos o son inválidos"})
		return
	}

	err = us_c.useCase.Run(idInt,&student)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	response := responses.NewResponseStudentUpdated(idInt, &student)

	ctx.JSON(http.StatusOK, response)

}