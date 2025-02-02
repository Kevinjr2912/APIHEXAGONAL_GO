package application

import (
	"api-hexagonal/career/domain/entities"
	"api-hexagonal/career/domain/repositories"
)

type UpdateCareer struct {
	db repositories.ICareer
}

func NewUpdateCareer(db repositories.ICareer) *UpdateCareer {
	return &UpdateCareer{db: db}
}

func (uc *UpdateCareer) Run(id int64, career *entities.Career) (err error) {
	return uc.db.UpdateCareer(id, career)
}