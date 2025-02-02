package controllers

import (
	"api-hexagonal/career/application"
	"api-hexagonal/career/infraestructure/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllCareersController struct {
	useCase *application.GetAllCareers
}

func NewGetAllCareersController(useCase *application.GetAllCareers) *GetAllCareersController {
	return &GetAllCareersController{useCase: useCase}
}

func (ga_c *GetAllCareersController) Run(ctx *gin.Context) {
	careers, err := ga_c.useCase.Run()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	response := responses.NewResponseGetAllCareers(careers)

	ctx.JSON(http.StatusOK, response)

}