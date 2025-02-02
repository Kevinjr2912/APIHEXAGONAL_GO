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
	getAllStudentsController := infraestructure.GetAllStudentsController().Run
	updateStudentController := infraestructure.UpdateStudentController().Run
	deleteStudentController := infraestructure.DeleteStudentController().Run

	// Definimos las rutas
	routes.POST("", createStudentController)
	routes.GET("", getAllStudentsController)
	routes.PUT("/:id", updateStudentController)
	routes.DELETE("/:id", deleteStudentController)
}