package controllers

import (
	"api-hexagonal/students/application"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type NotifyNewStudentController struct {
	useCase *application.GetAllStudents
	clients []chan string
	numStudents int
}

func NewNotifyStudentController(useCase *application.GetAllStudents) *NotifyNewStudentController {
	return &NotifyNewStudentController{useCase: useCase, clients: []chan string{}, numStudents: 0}
}

func (ns_c *NotifyNewStudentController) Run(ctx *gin.Context) {
	students, err := ns_c.useCase.Run()	

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// Lo que va a recibir el cliente por medio del canal (cadena de texto)
	channel := make(chan string)

	// Lista de clientes conectados
	ns_c.clients = append(ns_c.clients, channel)

	// Cantidad de estudiantes obtenidos de la base de datos
	currentNumStudents := len(*students)

	if ns_c.numStudents != 0 && ns_c.numStudents != currentNumStudents {
		for _, client := range ns_c.clients {
			go func(c chan string) {
				c <- "Se han registrado nuevos estudiantes"
			}(client)
		}
	}

	ns_c.numStudents = currentNumStudents
	ns_c.clients = []chan string{}

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Transfer-Encoding", "chunked")

	select {
	case msg := <-channel:
		ctx.JSON(http.StatusOK, gin.H{"Mensaje": msg})

	case <-time.After(15 * time.Second):
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Tiempo de espera agotado"})
	}

}


