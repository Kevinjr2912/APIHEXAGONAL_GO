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

// Verificar si hay nueva información de un estudiante
func CheckStudentUpdatesController() *controllers.CheckStudentUpdatesController {
	useCaseCheckStudentUpdates := application.NewCheckStudentUpdates(myql)

	return controllers.NewCheckStudentUpdatesController(useCaseCheckStudentUpdates)
}


// Encontrar un estudiante por medio del id
func FindByIdController() *controllers.FindByIdController {
	useCaseFindById := application.NewFindById(myql)

	return controllers.NewFindByIdController(useCaseFindById)
}


// Encontrar estudiantes por edad
func FindByAgeController() *controllers.FindByAgeController {
	useCaseFindByAge := application.NewFindByAge(myql)

	return controllers.NewFindByAgeController(useCaseFindByAge)
}

// Notificar a los clientes conectados cuando haya o hayan nuevos estudiantes
func NotifyNewStudentController() *controllers.NotifyNewStudentController {
	useCaseGetAllStudents := application.NewGetAllStudents(myql)

	return controllers.NewNotifyStudentController(useCaseGetAllStudents)
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