package infraestructure

import (
	"api-hexagonal/students/application"
	"api-hexagonal/students/infraestructure/controllers"
)

var (
	myql *MySQL
)

func Init() {
	myql = NewMySQL()
}

// Casos de uso

// Crear un estudiante
func CreateStudentController() *controllers.CreateStudentController {
	useCaseCreateStudent := application.NewCreateStudent(myql)

	return controllers.NewCreateStudentController(useCaseCreateStudent)
}

// Obtener todos los estudiantes
func GetAllStudentsController() *controllers.GetAllStudentsController {
	useCaseGetAllStudents := application.NewGetAllStudents(myql)

	return controllers.NewGetAllStudentsController(useCaseGetAllStudents)
}

// Actualizar la información de un estudiante
func UpdateStudentController() *controllers.UpdateStudentController {
	useCaseUpdateStudent := application.NewUpdateStudent(myql)

	return controllers.NewUpdateStudentController(useCaseUpdateStudent)
}

// Eliminar un estudiante
func DeleteStudentController() *controllers.DeleteStudentController {
	useCaseDeleteStudent := application.NewDeleteStudent(myql)

	return controllers.NewDeleteStudentController(useCaseDeleteStudent)
}