package responses

import (
	"api-hexagonal/career/domain/entities"
	"fmt"
)

type CareerAttributesUpdated struct {
	Name     string `json:"name"`
	Duration string `json:"duration"`
	Type     string `json:"typeCareer"`
}

type CareerLinksUpdated struct {
	Self string `json:"self"`
}

type ResponseCareerUpdated struct {
	Data struct {
		Type       string                  `json:"type"`
		Id         string                  `json:"id"`
		Attributes CareerAttributesUpdated `json:"attributes"`
		Links      CareerLinksUpdated      `json:"links"`
	} `json:"data"`
}

func NewResponseCareerUpdated(id int64, career *entities.Career) *ResponseCareerUpdated {
	return &ResponseCareerUpdated{
		Data: struct {
			Type       string                  `json:"type"`
			Id         string                  `json:"id"`
			Attributes CareerAttributesUpdated `json:"attributes"`
			Links      CareerLinksUpdated      `json:"links"`
		}{
			Type: "Careers",
			Id:   fmt.Sprintf("%d", id),
			Attributes: CareerAttributesUpdated{
				Name:        career.Name,
				Duration:         fmt.Sprintf("%d", career.Duration),
				Type: career.Type, 
			},
			Links: CareerLinksUpdated{
				Self: fmt.Sprintf("http://localhost:8080/careers/%d", id),
			},
		},
	}
}
