package repositories

import "api-hexagonal/career/domain/entities"

type ICareer interface {
	CreateCareer(career *entities.Career) (id int64, err error)
	GetAllCareers() (careers *[]entities.Career, err error)
	UpdateCareer(id int64, career *entities.Career) (err error)
	DeleteCareer(id int64) (err error)
}