package application

import (
	"api-hexagonal/students/domain/entities"
	"api-hexagonal/students/domain/repositories"
)

type FindByAge struct {
	db repositories.IStudent
}

func NewFindByAge(db repositories.IStudent) * FindByAge {
	return &FindByAge{db:db}
}

func (fba *FindByAge) Run(age uint8) (studentsArray *[]entities.Student, err error) {
	return fba.db.FindByAge(age)
}