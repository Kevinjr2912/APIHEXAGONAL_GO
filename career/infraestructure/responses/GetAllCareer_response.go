package responses

import (
	"api-hexagonal/career/domain/entities"
	"fmt"
)

type ResponseLinksStudents struct {
	Self string `json:"self"`
}

type ResponseDataCareer struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		Name     string `json:"name"`
		Duration string `json:"duration"`
		Type     string `json:"typeCareer"`
	} `json:"attributes"`
}

type ResponseGetAllStudents struct {
	Links ResponseLinksStudents  `json:"links"`
	Data  []ResponseDataCareer `json:"data"`
}

func NewResponseGetAllCareers(careers *[]entities.Career) *ResponseGetAllStudents {
	data := []ResponseDataCareer{}

	for i := 0; i < len(*careers); i++ {
		data = append(data, ResponseDataCareer{
			Type: "Careers",
			Id: fmt.Sprintf("%d", (*careers)[i].Id),
			Attributes: struct {
				Name     string `json:"name"`
				Duration string `json:"duration"`
				Type     string `json:"typeCareer"`
			}{
				Name: (*careers)[i].Name,
				Duration: fmt.Sprintf("%d",(*careers)[i].Duration),
				Type: (*careers)[i].Type,
			},
		})
	}

	return &ResponseGetAllStudents{
		Links: ResponseLinksStudents{
			Self: "http://localhost:8080/careers",
		},
		Data: data,
	}
}