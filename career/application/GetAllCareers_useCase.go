package application

import (
	"api-hexagonal/career/domain/entities"
	"api-hexagonal/career/domain/repositories"
)

type GetAllCareers struct {
	db repositories.ICareer
}

func NewGetAllCareers(db repositories.ICareer) *GetAllCareers {
	return &GetAllCareers{db: db}
}

func (gac *GetAllCareers) Run() (careers *[]entities.Career, err error) {
	return gac.db.GetAllCareers()
}