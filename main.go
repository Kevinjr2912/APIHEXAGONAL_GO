package main

import (
	"api-hexagonal/students/infraestructure"
	"api-hexagonal/students/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	infraestructure.Init()

	// Creamos el router
	r := gin.Default()

	routes.StudentRouter(r)

	// Levantamos el servidor
	r.Run()
}