package application

import "api-hexagonal/career/domain/repositories"

type DeleteCareer struct {
	db repositories.ICareer
}

func NewDeleteCareer(db repositories.ICareer) *DeleteCareer {
	return &DeleteCareer{db: db}
}

func (dc *DeleteCareer) Run(id int64) (err error) {
	return dc.db.DeleteCareer(id)
}