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
	getAllCareersController := infraestructure.GetAllCareersController().Run
	updateCareerController := infraestructure.UpdateCareerController().Run

	// Definimos las rutas
	routes.POST("", createCareerController)
	routes.GET("", getAllCareersController)
	routes.PUT("/:id", updateCareerController)
	
}