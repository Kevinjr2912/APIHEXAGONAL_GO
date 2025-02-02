package responses

import (
	"api-hexagonal/students/domain/entities"
	"fmt"
)

type StudentAttributesCreated struct {
	Name        string `json:"name"`
	Age        	string  `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
}

type StudentLinksCreated struct {
	Self string `json:"self"`
}

type ResponseStudentCreated struct {
	Data struct {
		Type       string                   `json:"type"`
		Id         string                   `json:"id"`
		Attributes StudentAttributesCreated `json:"attributes"`
		Links      StudentLinksCreated      `json:"links"`
	} `json:"data"`
}

func NewResponseStudentCreated(student *entities.Student) *ResponseStudentCreated {
	return &ResponseStudentCreated{
		Data: struct {
			Type       string                   `json:"type"`
			Id         string                   `json:"id"`
			Attributes StudentAttributesCreated `json:"attributes"`
			Links      StudentLinksCreated      `json:"links"`
		}{
			Type: "Students",
			Id:   fmt.Sprintf("%d", student.Id),
			Attributes: StudentAttributesCreated{
				Name:        student.Name,
				Age:         fmt.Sprintf("%d", student.Age),
				PhoneNumber: fmt.Sprintf("%d", student.PhoneNumber), 
			},
			Links: StudentLinksCreated{
				Self: fmt.Sprintf("http://localhost:8080/students/%d", student.Id),
			},
		},
	}
}
