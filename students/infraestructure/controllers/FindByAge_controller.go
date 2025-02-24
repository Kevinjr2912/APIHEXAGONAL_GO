package controllers

import (
	"api-hexagonal/students/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FindByAgeController struct {
	useCase *application.FindByAge
}

func NewFindByAgeController(useCase *application.FindByAge) *FindByAgeController {
	return &FindByAgeController{useCase: useCase}
}

func (fba_c *FindByAgeController) Run(ctx *gin.Context) {
	ageStr, exists := ctx.Params.Get("age")

	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "El parametro edad no se proporcionó"})
		return
	}

	ageUint, err := strconv.ParseInt(ageStr, 10, 8)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Formato invalido para la edad"})
		return
	}

	students , err := fba_c.useCase.Run(uint8(ageUint))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}


	ctx.JSON(http.StatusOK, students)
}