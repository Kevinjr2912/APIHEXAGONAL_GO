package application

import (
	"api-hexagonal/students/domain/entities"
	"api-hexagonal/students/domain/repositories"
)

type UpdateStudent struct {
	db repositories.IStudent
}

func NewUpdateStudent(db repositories.IStudent) *UpdateStudent {
	return &UpdateStudent{db: db}
}

func (up *UpdateStudent) Run(id int64, student *entities.Student) (err error) {
	return up.db.UpdateStudent(id, student)
}