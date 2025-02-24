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
	notifyNewCareerController := infraestructure.NotifyNewCareerController().Run
	updateCareerController := infraestructure.UpdateCareerController().Run
	deleteCareerController := infraestructure.DeleteCareerController().Run

	// Definimos las rutas
	routes.POST("", createCareerController)
	routes.GET("", getAllCareersController)
	routes.GET("/notification", notifyNewCareerController)
	routes.PUT("/:id", updateCareerController)
	routes.DELETE("/:id", deleteCareerController)
	
}