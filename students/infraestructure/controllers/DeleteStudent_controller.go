package controllers

import (
	"api-hexagonal/students/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteStudentController struct {
	useCase *application.DeleteStudent
}

func NewDeleteStudentController(usecase *application.DeleteStudent) *DeleteStudentController {
	return &DeleteStudentController{useCase: usecase}
}

func (ds_c *DeleteStudentController) Run(ctx *gin.Context) {
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

	err = ds_c.useCase.Run(idInt)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}