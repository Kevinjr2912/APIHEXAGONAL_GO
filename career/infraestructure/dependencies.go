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

// Notificar a todos los clientes de la existencia de nuevas carreras
func NotifyNewCareerController() *controllers.NotifyNewCareerController {
	useCaseGetAllCareers := application.NewGetAllCareers(mysql)

	return controllers.NewNotifyCareerController(useCaseGetAllCareers)
}

// Actualizar la información de una carrera
func UpdateCareerController() *controllers.UpdateCareerController {
	useCaseUpdateCareer := application.NewUpdateCareer(mysql)

	return controllers.NewUpdateCareerController(useCaseUpdateCareer)
}

// Eliminar una carrera
func DeleteCareerController() *controllers.DeleteCareerController {
	useCaseDeleteCareer := application.NewDeleteCareer(mysql)

	return controllers.NewDeleteCareerController(useCaseDeleteCareer)
}