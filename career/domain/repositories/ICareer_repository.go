package repositories

import "api-hexagonal/career/domain/entities"

type ICareer interface {
	CreateCareer(career *entities.Career) (id int64, err error)
}