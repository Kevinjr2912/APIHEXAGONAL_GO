package responses

import (
	"api-hexagonal/students/domain/entities"
	"fmt"
)

type StudentAttributesUpdated struct {
	Name        string `json:"name"`
	Age        	string  `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
}

type StudentLinksUpdated struct {
	Self string `json:"self"`
}

type ResponseStudentUpdated struct {
	Data struct {
		Type       string                   `json:"type"`
		Id         string                   `json:"id"`
		Attributes StudentAttributesUpdated `json:"attributes"`
		Links      StudentLinksUpdated      `json:"links"`
	} `json:"data"`
}

func NewResponseStudentUpdated(id int64, student *entities.Student) *ResponseStudentUpdated {
	return &ResponseStudentUpdated{
		Data: struct {
			Type       string                   `json:"type"`
			Id         string                   `json:"id"`
			Attributes StudentAttributesUpdated `json:"attributes"`
			Links      StudentLinksUpdated      `json:"links"`
		}{
			Type: "Students",
			Id:   fmt.Sprintf("%d", id),
			Attributes: StudentAttributesUpdated{
				Name:        student.Name,
				Age:         fmt.Sprintf("%d", student.Age),
				PhoneNumber: fmt.Sprintf("%d", student.PhoneNumber), 
			},
			Links: StudentLinksUpdated{
				Self: fmt.Sprintf("http://localhost:8080/students/%d", student.Id),
			},
		},
	}
}
