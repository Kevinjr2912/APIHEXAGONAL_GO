package repositories

import (
	"api-hexagonal/students/domain/entities"
	"time"
)

type IStudent interface {
	CreateStudent(student *entities.Student) (id int64, err error)
	GetAllStudents() (studentsArray *[]entities.Student, err error)
	GetStudentUpdatedAt(id int64) (time.Time, error)
	FindById(id int64) (student *entities.Student, err error)
	FindByAge(age uint8) (studentsArray *[]entities.Student, err error)
	UpdateStudent(id int64, student *entities.Student) (err error)
	DeleteStudent(id int64) (err error)
}