package main

import (
	studentInfra "api-hexagonal/students/infraestructure"
	careerInfra "api-hexagonal/career/infraestructure"
	studentRoutes	"api-hexagonal/students/infraestructure/routes"
	careerRoutes	"api-hexagonal/career/infraestructure/routes"
	
	"github.com/gin-gonic/gin"
)

func main() {
	studentInfra.Init()
	careerInfra.Init()

	// Creamos el router
	r := gin.Default()

	studentRoutes.StudentRouter(r)
	careerRoutes.CareerRouter(r)


	// Levantamos el servidor
	r.Run()
}