package application

import (
	"api-hexagonal/career/domain/entities"
	"api-hexagonal/career/domain/repositories"
)

type CreateCareer struct {
	db repositories.ICareer
}

func NewCreateCareer(db repositories.ICareer) *CreateCareer {
	return &CreateCareer{db: db}
}

func (cc *CreateCareer) Run(career *entities.Career) (id int64, err error) {
	return cc.db.CreateCareer(career)
}