package responses

import (
	"api-hexagonal/career/domain/entities"
	"fmt"
)

type CareerAttributesCreated struct {
	Name     string `json:"name"`
	Duration string `json:"duration"`
	Type     string `json:"typeCareer"`
}

type CareerLinksCreated struct {
	Self string `json:"self"`
}

type ResponseCareerCreated struct {
	Data struct {
		Type       string                  `json:"type"`
		Id         string                  `json:"id"`
		Attributes CareerAttributesCreated `json:"attributes"`
		Links      CareerLinksCreated      `json:"links"`
	} `json:"data"`
}

func NewResponseCareerCreated(career *entities.Career) *ResponseCareerCreated {
	return &ResponseCareerCreated{
		Data: struct{
			Type string `json:"type"`; 
			Id string `json:"id"`; 
			Attributes CareerAttributesCreated `json:"attributes"`; 
			Links CareerLinksCreated `json:"links"`
		}{
			Type: "Careers",
			Id: fmt.Sprintf("%d", career.Id),
			Attributes: CareerAttributesCreated{
				Name: career.Name,
				Duration: fmt.Sprintf("%d", career.Duration),
				Type: career.Type,
			},
			Links: CareerLinksCreated{
				Self: fmt.Sprintf("http://localhost:8080/careers/%d",career.Id),
			},
		},
	}
}