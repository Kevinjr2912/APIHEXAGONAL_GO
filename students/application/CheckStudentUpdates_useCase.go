package application

import (
	"api-hexagonal/students/domain/repositories"
	"time"
)

type CheckStudentUpdates struct {
	db repositories.IStudent
}

func NewCheckStudentUpdates(db repositories.IStudent) *CheckStudentUpdates {
	return &CheckStudentUpdates{db: db}
}

func (csu *CheckStudentUpdates) Run(id int64) (time.Time, error) {
	return csu.db.GetStudentUpdatedAt(id)
}
