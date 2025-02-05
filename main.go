package main

import (
	careerInfra "api-hexagonal/career/infraestructure"
	careerRoutes "api-hexagonal/career/infraestructure/routes"
	"api-hexagonal/core"
	studentInfra "api-hexagonal/students/infraestructure"
	studentRoutes "api-hexagonal/students/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	studentInfra.Init()
	careerInfra.Init()

	// Creamos el router
	r := gin.Default()

	// CORS
	core.InitCORS(r)

	studentRoutes.StudentRouter(r)
	careerRoutes.CareerRouter(r)

	// Levantamos el servidor
	r.Run()
}