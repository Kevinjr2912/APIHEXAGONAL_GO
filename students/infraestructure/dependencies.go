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
	useCaseCreateStudent := application.NewCreateStudent(myql)

	return controllers.NewCreateStudentController(useCaseCreateStudent)
}

func GetAllStudentsController() *controllers.GetAllStudentsController {
	useCaseGetAllStudents := application.NewGetAllStudents(myql)

	return controllers.NewGetAllStudentsController(useCaseGetAllStudents)
}