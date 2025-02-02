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

// Crear una carrera
func CreateCareerController() *controllers.CreateCareerController {
	useCaseCreateCareer := application.NewCreateCareer(mysql)

	return controllers.NewCreateCareerController(useCaseCreateCareer)
}

// Obtener todas las carreras
func GetAllCareersController() *controllers.GetAllCareersController  {
	useCaseGetAllCareers := application.NewGetAllCareers(mysql)

	return controllers.NewGetAllCareersController(useCaseGetAllCareers)
}
