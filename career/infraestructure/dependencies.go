package infraestructure

import (
	"api-hexagonal/career/application"
	"api-hexagonal/career/infraestructure/controllers"
)

var (
	mysql *MySQL
)

func Init() {
	mysql = NewMySQL()
}

// Casos de uso
func CreateCareerController() *controllers.CreateCareerController {
	useCase := application.NewCreateCareer(mysql)

	return controllers.NewCreateCareerController(useCase)
}