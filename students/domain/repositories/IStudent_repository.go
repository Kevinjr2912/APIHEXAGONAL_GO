package repositories

import "api-hexagonal/students/domain/entities"

type IStudent interface {
	CreateStudent(student *entities.Student) (id int64, err error)
	GetAllStudents() (studentsArray *[]entities.Student, err error)
	// UpdateStudent()
	// DeleteStudent()
}