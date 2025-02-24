package controllers

import (
	"api-hexagonal/career/application"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type NotifyNewCareerController struct {
	useCase *application.GetAllCareers
	clients []chan string
	numCareers int
}

func NewNotifyCareerController(useCase *application.GetAllCareers) *NotifyNewCareerController {
	return &NotifyNewCareerController{useCase: useCase}
}

func (nc_c *NotifyNewCareerController) Run(ctx *gin.Context) {
	careers, err := nc_c.useCase.Run()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	channel := make(chan string)

	nc_c.clients = append(nc_c.clients, channel)

	currentNumCareers := len(*careers)

	if nc_c.numCareers != 0 && nc_c.numCareers != currentNumCareers {
		for _, career := range nc_c.clients {
			go func (c chan string) {
				c <- "Nuevas carreras, dale un vistazo"
			}(career)
		}	
	}

	nc_c.numCareers = currentNumCareers
	nc_c.clients = []chan string{}

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Transfer-Encoding", "chunked")

	select {
	case msg := <-channel:
		ctx.JSON(http.StatusOK, gin.H{"Mensaje": msg})

	case <-time.After(15 * time.Second):
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Tiempo de espera agotado"})
	}
}