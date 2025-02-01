package routes

import (
	"api-hexagonal/students/infraestructure"

	"github.com/gin-gonic/gin"
)

func StudentRouter(router *gin.Engine) {
	// Definimos el grupo de router
	routes := router.Group("/students") 

	// Instanciamos los controladores
	createStudentController := infraestructure.CreateStudentController().Run

	// Definimos las rutas
	routes.POST("", createStudentController)
}