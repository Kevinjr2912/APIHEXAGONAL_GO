package responses

import (
	"api-hexagonal/students/domain/entities"
	"fmt"
)

type ResponseLinkStudents struct {
	Links string `json:"links"`
}

type ResponseDataStudents struct {
	Type       string `json:"type"`
	Id         string  `json:"id"`
	Attributes struct {
		Name        string `json:"name"`
		Age         string  `json:"age"`
		PhoneNumber string `json:"phoneNumber"`
	} `json:"attributes"`
}

type ResponseGetAllStudents struct {
	Links ResponseLinkStudents   `json:"links"`
	Data  []ResponseDataStudents `json:"data"`
}

func NewResponseGetAllStudents(students *[]entities.Student) *ResponseGetAllStudents {
	data := []ResponseDataStudents{}

	for i := 0; i < len(*students); i++ {
		data = append(data, ResponseDataStudents{
			Type: "Students",
			Id:   fmt.Sprintf("%d", (*students)[i].Id),
			Attributes: struct {
				Name        string  `json:"name"`
				Age         string  `json:"age"`
				PhoneNumber string  `json:"phoneNumber"`
			}{
				Name:        fmt.Sprintf((*students)[i].Name),
				Age:         fmt.Sprintf("%d", (*students)[i].Age),
				PhoneNumber: fmt.Sprintf("%d", (*students)[i].PhoneNumber),
			},
		})
	}

	return &ResponseGetAllStudents{
		Links: ResponseLinkStudents{
			Links: "http://localhost:8080/students/getAllStudents",
		},
		Data: data,
	}
}
