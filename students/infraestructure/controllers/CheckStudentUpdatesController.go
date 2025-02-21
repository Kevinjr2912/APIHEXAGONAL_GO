package controllers

import (
	"api-hexagonal/students/application"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckStudentUpdatesController struct {
	useCase *application.CheckStudentUpdates
}

func NewCheckStudentUpdatesController(useCase *application.CheckStudentUpdates) *CheckStudentUpdatesController {
	return &CheckStudentUpdatesController{useCase: useCase}
}

func (csuc *CheckStudentUpdatesController) Run(ctx *gin.Context) {
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

	sinceStr := ctx.DefaultQuery("since", "")
	var lastCheck time.Time

	if sinceStr != "" {
		lastCheck, err = time.Parse(time.RFC3339, sinceStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Formato de timestamp inválido"})
			return
		}
	} else {
		lastCheck = time.Now()
	}

	updateAt, err := csuc.useCase.Run(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	hasUpdates := updateAt.After(lastCheck)

	if hasUpdates {
		ctx.JSON(http.StatusOK, gin.H{"updated": true})
	} else {
		ctx.JSON(http.StatusNotModified, gin.H{"updated": false})
	}
}
