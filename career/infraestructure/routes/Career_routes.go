package routes

import (
	"api-hexagonal/career/infraestructure"

	"github.com/gin-gonic/gin"
)

func CareerRouter(router *gin.Engine) {
	// Definimos el grupo de router
	routes := router.Group("/careers")

	// Instanciamos los controladores
	createCareerController := infraestructure.CreateCareerController().Run

	// Definimos las rutas
	routes.POST("", createCareerController)
	
}