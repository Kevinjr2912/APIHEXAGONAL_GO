package responses

import (
	"api-hexagonal/students/domain/entities"
	"fmt"
)

type StudentAttributes struct {
	Name        string `json:"name"`
	Age        	string  `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
}

type StudentLinks struct {
	Self string `json:"self"`
}

type ResponseCreatedStudent struct {
	Data struct {
		Type       string            `json:"type"`
		Id         string            `json:"id"`
		Attributes StudentAttributes `json:"attributes"`
		Links      StudentLinks      `json:"links"`
	} `json:"data"`
}

func NewResponseCreatedStudent(student *entities.Student) *ResponseCreatedStudent {
	return &ResponseCreatedStudent{
		Data: struct {
			Type       string            `json:"type"`
			Id         string            `json:"id"`
			Attributes StudentAttributes `json:"attributes"`
			Links      StudentLinks      `json:"links"`
		}{
			Type: "Students",
			Id:   fmt.Sprintf("%d", student.Id),
			Attributes: StudentAttributes{
				Name:        student.Name,
				Age:         fmt.Sprintf("%d", student.Age),
				PhoneNumber: fmt.Sprintf("%d", student.PhoneNumber), 
			},
			Links: StudentLinks{
				Self: fmt.Sprintf("http://localhost:8080/students/%d", student.Id),
			},
		},
	}
}
