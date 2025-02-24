package responses

import (
	"api-hexagonal/students/domain/entities"
	"fmt"
)

type StudentAttributesFound struct {
	Name        string `json:"name"`
	Age         string `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
}

type StudentLinksFound struct {
	Self string `json:"self"`
}

type ResponseStudentFound struct {
	Data struct {
		Type       string                   `json:"type"`
		Id         string                   `json:"id"`
		Attributes StudentAttributesFound   `json:"attributes"`
		Links      StudentLinksFound        `json:"links"`
	} `json:"data"`
}

func NewResponseStudentFound(student *entities.Student) *ResponseStudentFound {
	return &ResponseStudentFound{
		Data: struct {
			Type       string                   `json:"type"`
			Id         string                   `json:"id"`
			Attributes StudentAttributesFound   `json:"attributes"`
			Links      StudentLinksFound        `json:"links"`
		}{
			Type: "Students",
			Id:   fmt.Sprintf("%d", student.Id),
			Attributes: StudentAttributesFound{
				Name:        student.Name,
				Age:         fmt.Sprintf("%d", student.Age),
				PhoneNumber: fmt.Sprintf("%d", student.PhoneNumber),
			},
			Links: StudentLinksFound{
				Self: fmt.Sprintf("http://localhost:8080/students/%d", student.Id),
			},
		},
	}
}
