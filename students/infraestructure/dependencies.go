package infraestructure

import (
	"api-hexagonal/students/application"
	"api-hexagonal/students/infraestructure/controllers"
)

var myql *MySQL

func Init() {
	myql = NewMySQL()
}

// Casos de uso
func CreateStudentController() *controllers.CreateStudentController {
	useCaseCreateStuden := application.NewCreateStudent(myql)

	return controllers.NewCreateStudentController(useCaseCreateStuden)
}